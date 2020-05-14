package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/one-go/xconf/api"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	prefix   = "/xconf/"
	pfxConf  = prefix + "xconf/"
	pfxSpace = pfxConf + "spaces/"
	pfxGroup = pfxConf + "group/"
)

type Console struct {
	cli    *clientv3.Client
	logger *zap.Logger
}

func newConsoleServer(logger *zap.Logger, client *clientv3.Client) *Console {
	console := &Console{
		cli:    client,
		logger: logger,
	}
	return console
}

func (c *Console) CreateNamespace(ctx context.Context, in *pb.CreateNamespaceRequest) (*pb.Namespace, error) {
	_, err := c.cli.Put(ctx, pfxSpace+in.GetSpace().Name, "")
	if err != nil {
		return nil, err
	}
	return in.Space, nil
}

func (c *Console) ListNamespaces(ctx context.Context, in *empty.Empty) (*pb.ListNamespacesResponse, error) {
	res, err := c.cli.Get(ctx, pfxSpace, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListNamespacesResponse)
	for _, kv := range res.Kvs {
		space := strings.TrimPrefix(string(kv.Key), pfxSpace)
		resp.Spaces = append(resp.Spaces, &pb.Namespace{Name: space})
	}
	return resp, nil
}

func (c *Console) CreateGroup(ctx context.Context, in *pb.CreateGroupRequest) (*pb.Group, error) {
	_, err := c.cli.Put(ctx, pfxGroup+in.Namespace+"/"+in.GetGroup().Name, "")
	if err != nil {
		return nil, err
	}
	return in.Group, nil
}

func (c *Console) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	pfx := pfxGroup + in.Namespace + "/"
	res, err := c.cli.Get(ctx, pfx, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListGroupsResponse)
	for _, kv := range res.Kvs {
		group := strings.TrimPrefix(string(kv.Key), pfx)
		resp.Groups = append(resp.Groups, &pb.Group{Name: group})
	}
	return resp, nil
}

func (c *Console) CreateConfig(ctx context.Context, in *pb.CreateConfigRequest) (*pb.Config, error) {
	if in.Parent == "" || in.Config == nil || in.Config.Id == "" {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}
	name := prefix + in.Parent + "/" + in.Config.Id
	res, err := c.cli.Get(ctx, name+".metadata")
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return nil, status.Error(codes.AlreadyExists, "config already exists")
	}
	if in.Config.Meta == nil {
		in.Config.Meta = new(pb.ConfigMeta)
	}
	in.Config.Meta.Ctime = ptypes.TimestampNow()
	if in.Config.Meta.Version == "" {
		// in.Config.Meta.Version = strconv.FormatInt(in.Config.Meta.Ctime.GetSeconds(), 10)
	}
	if err := c.putConfig(ctx, name, in.Config); err != nil {
		return nil, err
	}
	return in.Config, nil
}

func (c *Console) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.Config, error) {
	name := prefix + in.Parent + "/" + in.Config.Id
	config, err := c.getConfig(ctx, name, in.Config.Id)
	if err != nil {
		return nil, err
	}
	for _, path := range in.UpdateMask.GetPaths() {
		switch path {
		case "config.content":
			config.Content = in.Config.Content
		case "config.meta.version":
			config.Meta.Version = in.Config.Meta.Version
		case "config.meta.canary":
			config.Meta.Canary = in.Config.Meta.Canary
		case "config.meta.comment":
			config.Meta.Comment = in.Config.Meta.Comment
		}
	}
	config.Meta.Mtime = ptypes.TimestampNow()
	if err = c.putConfig(ctx, name, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Console) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.Config, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "config name empty")
	}
	name := prefix + in.GetName()
	return c.getConfig(ctx, name, in.Name)
}

func (c *Console) DeleteConfig(ctx context.Context, in *pb.DeleteConfigRequest) (*empty.Empty, error) {
	name := prefix + in.GetName()
	if _, err := c.cli.Delete(ctx, name+".metadata"); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *Console) ListConfigs(ctx context.Context, in *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	pfx := prefix + in.Parent + "/"
	res, err := c.cli.Get(ctx, pfx, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListConfigsResponse)
	for _, kv := range res.Kvs {
		name := string(kv.Key)
		if config, err := c.getConfig(ctx, name, strings.TrimPrefix(name, pfx)); err == nil {
			resp.Configs = append(resp.Configs, config)
		}
	}
	return resp, nil
}

func (c *Console) getConfig(ctx context.Context, name, id string) (*pb.Config, error) {
	content, err := Get(ctx, c.cli, name)
	if err != nil {
		return nil, err
	}
	meta := new(pb.ConfigMeta)
	if err = GetObject(ctx, c.cli, name+".metadata", meta); err != nil {
		return nil, err
	}
	config := &pb.Config{
		Id:      id,
		Meta:    meta,
		Content: string(content),
	}
	return config, nil
}

func (c *Console) putConfig(ctx context.Context, name string, in *pb.Config) error {
	meta, _ := json.Marshal(in.GetMeta())
	if _, err := c.cli.Put(ctx, name, in.Content); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if _, err := c.cli.Put(ctx, name+".metadata", string(meta)); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

// etcd wrap api
func Get(ctx context.Context, c *clientv3.Client, k string) ([]byte, error) {
	res, err := c.Get(ctx, k)
	if err != nil {
		return nil, status.Error(codes.Internal, "etcd error")
	}
	if res.Count == 0 {
		return nil, status.Error(codes.NotFound, "config not found")
	}
	return res.Kvs[0].Value, nil
}

func GetObject(ctx context.Context, c *clientv3.Client, k string, v interface{}) error {
	data, err := Get(ctx, c, k)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}