package console

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/one-go/xconf/api"
	"github.com/one-go/xconf/internal/config"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Console struct {
	cli    *config.Client
	logger *zap.Logger
}

func NewServer(logger *zap.Logger, db *clientv3.Client) *Console {
	console := &Console{
		cli:    config.New(db),
		logger: logger,
	}
	return console
}

func (c *Console) CreateNamespace(ctx context.Context, in *pb.CreateNamespaceRequest) (*pb.Namespace, error) {
	if err := c.cli.CreateNamespace(ctx, in.GetSpace().Name); err != nil {
		return nil, err
	}
	return in.Space, nil
}

func (c *Console) ListNamespaces(ctx context.Context, in *empty.Empty) (*pb.ListNamespacesResponse, error) {
	spaces, err := c.cli.ListNamespaces(ctx)
	if err != nil {
		return nil, err
	}
	resp := new(pb.ListNamespacesResponse)
	for _, space := range spaces {
		resp.Spaces = append(resp.Spaces, &pb.Namespace{Name: space})
	}
	return resp, nil
}

func (c *Console) CreateGroup(ctx context.Context, in *pb.CreateGroupRequest) (*pb.Group, error) {
	return in.Group, nil
}

func (c *Console) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	groups, err := c.cli.ListGroups(ctx, in.Namespace)
	if err != nil {
		return nil, err
	}
	resp := new(pb.ListGroupsResponse)
	for _, group := range groups {
		resp.Groups = append(resp.Groups, &pb.Group{Name: group})
	}
	return resp, nil
}

func (c *Console) CreateConfig(ctx context.Context, in *pb.CreateConfigRequest) (*pb.Config, error) {
	if in.Parent == "" || in.Config == nil || in.Config.Id == "" {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}
	if in.Config.Meta == nil {
		in.Config.Meta = new(pb.ConfigMeta)
	}
	in.Config.Meta.Ctime = ptypes.TimestampNow()
	if in.Config.Meta.Version == "" {
		// in.Config.Meta.Version = strconv.FormatInt(in.Config.Meta.Ctime.GetSeconds(), 10)
	}
	if err := c.cli.CreateConfig(ctx, in.Parent, in.Config); err != nil {
		return nil, err
	}
	return in.Config, nil
}

func (c *Console) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.Config, error) {
	config, err := c.cli.GetConfig(ctx, in.Parent, in.Config.Id)
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
	if err = c.cli.PutConfig(ctx, in.Parent, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Console) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.Config, error) {
	if in.Parent == "" || in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "config name empty")
	}
	return c.cli.GetConfig(ctx, in.Parent, in.Id)
}

func (c *Console) DeleteConfig(ctx context.Context, in *pb.DeleteConfigRequest) (*empty.Empty, error) {
	if in.Parent == "" || in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "config name empty")
	}
	if err := c.cli.DeleteConfig(ctx, in.Parent, in.Id); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *Console) ListConfigs(ctx context.Context, in *pb.ListConfigsRequest) (*pb.ListConfigsResponse, error) {
	configs, err := c.cli.ListConfigs(ctx, in.Parent)
	if err != nil {
		return nil, err
	}

	resp := new(pb.ListConfigsResponse)
	for _, config := range configs {
		resp.Configs = append(resp.Configs, config)
	}
	return resp, nil
}
