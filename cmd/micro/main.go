package main

import (
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"
	"github.com/micro/micro/cmd"
)

func init() {
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()
}

func main() {
	cmd.Init()
}
