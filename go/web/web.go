package web

import (
	"github.com/micro/go-micro"
	cli "github.com/micro/go-plugins/client/grpc"
	"github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"
	"github.com/micro/go-web"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"github.com/micro/go-plugins/selector/static"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// setup
	c := cli.NewClient()
	s := srv.NewServer()
	k := kubernetes.NewRegistry()
	st := static.NewSelector()

	// create new service
	service := micro.NewService(
		micro.Server(s),
		micro.Client(c),
		micro.Registry(k),
		micro.Selector(st),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
