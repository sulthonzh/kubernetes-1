package main

import (
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"
	"github.com/micro/micro/cmd"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/micro/go-plugins/selector/static"

	// disable namespace by default
	"github.com/micro/micro/api"
)

func main() {
	// disable namespace
	api.Namespace = ""

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")

	// setup client/server
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// init command
	cmd.Init()
}
