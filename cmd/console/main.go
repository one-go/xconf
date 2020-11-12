//go:generate protoc -I api --go_out=plugins=grpc:api --js_out=import_style=commonjs:web --grpc-web_out=mode=grpcwebtext:web api/xconf.proto

package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"strings"

	pb "github.com/one-go/xconf/api"
	"github.com/one-go/xconf/internal/console"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host       = flag.String("l", ":8900", "listen address")
	staticHost = flag.String("s", "", "listen static files")
	staticDir  = flag.String("dir", "web", "static dir")
	endpoints  = flag.String("h", os.Getenv("ETCD_ENDPOINTS"), "etcd endpoints")
	username   = flag.String("u", os.Getenv("ETCD_USER"), "etcd username")
	passwd     = flag.String("p", os.Getenv("ETCD_PASSWD"), "etcd password")
)

func main() {
	flag.Parse()
	log, _ := zap.NewProduction()
	defer log.Sync()

	if *staticHost != "" {
		go func() {
			if err := http.ListenAndServe(*staticHost, http.FileServer(http.Dir(*staticDir))); err != nil {
				log.Fatal("file server failed", zap.String("host", *staticHost), zap.String("dir", *staticDir), zap.Error(err))
			}
		}()
	}

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatal("failed to listen", zap.String("host", *host))
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(*endpoints, ","),
		Username:  *username,
		Password:  *passwd,
	})
	if err != nil {
		log.Fatal("failed to connect etcd", zap.String("host", *host), zap.Error(err))
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	s := console.NewServer(log, cli)
	pb.RegisterXconfServer(grpcServer, s)
	log.Info("listening on", zap.String("host", *host))
	grpcServer.Serve(lis)
}
