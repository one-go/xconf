//go:generate protoc -I api --go_out=plugins=grpc:api --js_out=import_style=commonjs:api --grpc-web_out=mode=grpcwebtext:api api/xconf.proto

package main

import (
	"flag"
	"log"
	"net"
	"strings"

	pb "github.com/one-go/xconf/console/api"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Interval      int    `toml:"interval"`
	SecretKeyring string `toml:"secret_keyring"`
	SRVDomain     string `toml:"srv_domain"`
	SRVRecord     string `toml:"srv_record"`
	LogLevel      string `toml:"log-level"`
	Watch         bool   `toml:"watch"`
	PrintVersion  bool
	ConfigFile    string
	OneTime       bool
}

var config Config

func main() {
	flag.Parse()
	host := flag.String("host", "127.0.0.1:2378", "listen address")
	endpoints := flag.String("endpoints", "127.0.0.1:2379", "etcd endpoints")
	username := flag.String("user", "", "etcd username")
	passwd := flag.String("passwd", "", "etcd password")

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(*endpoints, ","),
		Username:  *username,
		Password:  *passwd,
	})
	if err != nil {
		log.Fatalf("falied to connect etcd: %+v", err)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterXconfServer(grpcServer, newConsoleServer(logger, cli))
	grpcServer.Serve(lis)
}
