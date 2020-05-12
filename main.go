//go:generate protoc -I api --go_out=plugins=grpc:api --js_out=import_style=commonjs:api --grpc-web_out=mode=grpcwebtext:api api/xconf.proto

package main

import (
	"flag"
	"net"
	"os"
	"strings"

	pb "github.com/one-go/xconf/api"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	host := flag.String("l", "127.0.0.1:8900", "listen address")
	endpoints := flag.String("h", os.Getenv("ETCD_ENDPOINTS"), "etcd endpoints")
	username := flag.String("u", os.Getenv("ETCD_USER"), "etcd username")
	passwd := flag.String("p", os.Getenv("ETCD_PASSWD"), "etcd password")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		logger.Fatal("failed to listen", zap.String("host", *host))
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(*endpoints, ","),
		Username:  *username,
		Password:  *passwd,
	})
	if err != nil {
		logger.Fatal("failed to connect etcd", zap.String("host", *host), zap.Error(err))
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	s := newConsoleServer(logger, cli)
	pb.RegisterXconfServer(grpcServer, s)
	logger.Info("listening on", zap.String("host", *host))
	grpcServer.Serve(lis)
}
