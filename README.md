# Micro on Kubernetes

This repo provides the config to run Micro on Kubernetes

Services make use of the [kubernetes registry](https://github.com/micro/go-plugins/tree/master/registry/kubernetes) 
plugin so there's zero external dependency for service discovery.

## What's in the repo?

- cmd/micro - a preinitialised version of the micro toolkit with the kubernetes registry
	* go get github.com/micro/kubernetes/cmd/micro
- config/micro - API, Web UI and Sidecar (spins up GCE Load Balancers)
- config/services - Some example micro services

## Getting Started

Here's the steps I took to get started.

### Run Kubernetes

GKE is the easiest way to run a managed kubernetes cluster. What's even better? $300 free credit for 60 days.

1. Get yourself a [Free Trial](https://cloud.google.com/free-trial/) of Google Container Engine
2. Visit the [Quickstart](https://cloud.google.com/container-engine/docs/quickstart) guide to create a cluster

### Run Micro

We provide some predefined k8s configs with prebuilt micro images.

Make sure kubectl is in your path

```shell
./run.sh start
```

### Micro Toolkit

Get a preinitialised version of the micro toolkit with the kubernetes registry

```
go get github.com/micro/kubernetes/cmd/micro
```

### Build Yourself

```
go get github.com/micro/micro
```

Create a plugins.go file
```go
package main

import _ "github.com/micro/go-plugins/registry/kubernetes"
```

Build binary
```shell
// For local use
go build -i -o micro ./main.go ./plugins.go
```

Flag usage of plugins
```shell
micro --registry=kubernetes
```

### With Go Micro

Use the plugin with go-micro

```go
package main

import (
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("my.service"),
	)
	service.Init()
	service.Run()
}
```

```
go run main.go --registry=kubernetes
```

