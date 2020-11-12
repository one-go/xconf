package config

import (
	"context"
	"encoding/json"
	"strings"

	pb "github.com/one-go/xconf/api"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	prefix   = "/xconf/"
	pfxConf  = prefix + "xconf/"
	pfxSpace = pfxConf + "spaces/"
	pfxGroup = pfxConf + "group/"
)

type OnChange func(*pb.Config)

type Client struct {
	db *clientv3.Client
}

func New(db *clientv3.Client) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) CreateNamespace(ctx context.Context, ns string) error {
	_, err := c.db.Put(ctx, pfxSpace+ns, "")
	return err
}

func (c *Client) ListNamespaces(ctx context.Context) ([]string, error) {
	res, err := c.db.Get(ctx, pfxSpace, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var spaces []string
	for _, kv := range res.Kvs {
		space := strings.TrimPrefix(string(kv.Key), pfxSpace)
		spaces = append(spaces, space)
	}
	return spaces, nil
}

func (c *Client) CreateGroup(ctx context.Context, ns, group string) error {
	_, err := c.db.Put(ctx, pfxGroup+ns+"/"+group, "")
	return err
}

func (c *Client) ListGroups(ctx context.Context, ns string) ([]string, error) {
	pfx := pfxGroup + ns + "/"
	res, err := c.db.Get(ctx, pfx, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var groups []string
	for _, kv := range res.Kvs {
		group := strings.TrimPrefix(string(kv.Key), pfx)
		groups = append(groups, group)
	}
	return groups, nil
}

func (c *Client) CreateConfig(ctx context.Context, parent string, config *pb.Config) error {
	name := prefix + parent + "/" + config.Id
	res, err := c.db.Get(ctx, name+".metadata")
	if err != nil {
		return err
	}
	if res.Count > 0 {
		return status.Error(codes.AlreadyExists, "config already exists")
	}
	if err := c.PutConfig(ctx, name, config); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetConfig(ctx context.Context, parent, id string) (*pb.Config, error) {
	name := prefix + parent + "/" + id
	content, err := Get(ctx, c.db, name)
	if err != nil {
		return nil, err
	}
	meta := new(pb.ConfigMeta)
	if err = GetObject(ctx, c.db, name+".metadata", meta); err != nil {
		return nil, err
	}
	config := &pb.Config{
		Id:      id,
		Meta:    meta,
		Content: string(content),
	}
	return config, nil
}

func (c *Client) PutConfig(ctx context.Context, parent string, in *pb.Config) error {
	name := prefix + parent + "/" + in.Id
	return c.putConfig(ctx, name, in)
}

func (c *Client) DeleteConfig(ctx context.Context, parent, id string) error {
	name := prefix + parent + "/" + id
	_, err := c.db.Delete(ctx, name+".metadata")
	return err
}

func (c *Client) ListConfigs(ctx context.Context, parent string) ([]*pb.Config, error) {
	pfx := prefix + parent + "/"
	res, err := c.db.Get(ctx, pfx, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	var configs []*pb.Config
	for _, kv := range res.Kvs {
		name := string(kv.Key)
		if !strings.HasSuffix(name, ".metadata") {
			id := name[len(pfx):]
			if config, err := c.GetConfig(ctx, parent, id); err == nil {
				configs = append(configs, config)
			}
		}
	}
	return configs, nil
}

func (c *Client) putConfig(ctx context.Context, name string, in *pb.Config) error {
	meta, _ := json.Marshal(in.GetMeta())
	if _, err := c.db.Put(ctx, name, in.Content); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if _, err := c.db.Put(ctx, name+".metadata", string(meta)); err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (c *Client) WatchConfig(ctx context.Context, parent, id string, h OnChange) {
	name := prefix + parent + "/" + id
	ch := c.db.Watch(ctx, name+".metadata")
	for {
		select {
		case wr := <-ch:
			for _, ev := range wr.Events {
				meta := new(pb.ConfigMeta)
				if err := json.Unmarshal(ev.Kv.Value, &meta); err != nil {
					continue
				}
				content, err := Get(ctx, c.db, name)
				if err != nil {
					continue
				}
				config := &pb.Config{
					Id:      id,
					Meta:    meta,
					Content: string(content),
				}
				h(config)
			}
		case <-ctx.Done():
			break
		}
	}
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
