//go:generate protoc -I api --go_out=plugins=grpc:api --js_out=import_style=commonjs:api --grpc-web_out=mode=grpcwebtext:api api/xconf.proto

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/one-go/xconf/api"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	host := flag.String("listener", "127.0.0.1:8900", "listen address")
	endpoints := flag.String("endpoints", os.Getenv("ETCD_ENDPOINTS"), "etcd endpoints")
	username := flag.String("user", os.Getenv("ETCD_USER"), "etcd username")
	passwd := flag.String("passwd", os.Getenv("ETCD_PASSWD"), "etcd password")
	flag.Parse()

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
	reflection.Register(grpcServer)
	s := newConsoleServer(logger, cli)
	pb.RegisterXconfServer(grpcServer, s)
	grpcServer.Serve(lis)
}
