package main

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/one-go/xconf/console/api"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pfxConf  = "/xconf/xconf/"
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

func (c *Console) CreateNamespace(ctx context.Context, in *pb.CreateNamespaceRequest) (*empty.Empty, error) {
	_, err := c.cli.Put(ctx, pfxSpace+in.Namespace, "")
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *Console) ListNamespaces(ctx context.Context, in *empty.Empty) (*pb.ListNamespacesResponse, error) {
	res, err := c.cli.Get(ctx, pfxSpace, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListNamespacesResponse)
	for _, kv := range res.Kvs {
		resp.Namespaces = append(resp.Namespaces, string(kv.Key))
	}
	return resp, nil
}

func (c *Console) CreateGroup(ctx context.Context, in *pb.CreateGroupRequest) (*empty.Empty, error) {
	_, err := c.cli.Put(ctx, pfxGroup+in.Name, "")
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *Console) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	res, err := c.cli.Get(ctx, pfxGroup+in.Namespace, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListGroupsResponse)
	for _, kv := range res.Kvs {
		resp.Names = append(resp.Names, string(kv.Key))
	}
	return resp, nil
}

func (c *Console) CreateConfig(ctx context.Context, in *pb.Config) (*pb.Config, error) {
	if _, err := Get(ctx, c.cli, c.MetaName(in.Name)); err != nil {
		return nil, err
	}
	in.Meta.Ctime = ptypes.TimestampNow()
	if err := c.putConfig(ctx, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (c *Console) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.Config, error) {
	config, err := c.getConfig(ctx, in.GetConfig().Name)
	if err != nil {
		return nil, err
	}
	for _, path := range in.UpdateMask.GetPaths() {
		switch path {
		case "config.content":
			config.Content = in.Config.Content
		case "config.meta.canary":
			config.Meta.Canary = in.Config.Meta.Canary
		case "config.meta.comment":
			config.Meta.Comment = in.Config.Meta.Comment
		}
	}
	config.Meta.Mtime = ptypes.TimestampNow()
	if err = c.putConfig(ctx, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Console) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.Config, error) {
	return c.getConfig(ctx, in.GetName())
}

func (c *Console) getConfig(ctx context.Context, name string) (*pb.Config, error) {
	content, err := Get(ctx, c.cli, c.Name(name))
	if err != nil {
		return nil, err
	}
	meta := new(pb.ConfigMeta)
	if err = GetObject(ctx, c.cli, c.MetaName(name), meta); err != nil {
		return nil, err
	}
	config := &pb.Config{
		Name:    name,
		Meta:    meta,
		Content: string(content),
	}
	return config, nil
}

func (c *Console) DeleteConfig(ctx context.Context, in *pb.DeleteConfigRequest) (*empty.Empty, error) {
	if _, err := c.cli.Delete(ctx, c.MetaName(in.Name)); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *Console) ListConfigs(ctx context.Context, in *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	return nil, nil
}

func (c *Console) Name(name string) string {
	return "/xconf/" + name
}

func (c *Console) MetaName(name string) string {
	return c.Name(name) + ".metadata"
}

func (c *Console) putConfig(ctx context.Context, in *pb.Config) error {
	meta, _ := json.Marshal(in.GetMeta)
	if _, err := c.cli.Put(ctx, c.Name(in.Name), in.Content); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if _, err := c.cli.Put(ctx, c.MetaName(in.Name), string(meta)); err != nil {
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
