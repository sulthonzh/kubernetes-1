# Micro on Kubernetes

Micro on Kubernetes is a pre-tuned version of micro/go-micro for k8s.

## Features

- Kubernetes registry
- gRPC client and server
- caching selector
- pre-initialised toolkit

## Contents

- go/micro - pre-initialised go-micro
- cmd/micro - pre-initialised micro toolkit
- cmd/health - a healthchecking sidecar 

Example configs:

- config/micro - API, Web UI and Sidecar (spins up GCE Load Balancers)
- config/services - Some example micro services

## Getting Started

- [Writing a Service](#writing-a-service)
- [Install Micro](#install-micro)
- [Healthchecking Sidecar](#healthchecking-sidecar)

### Writing a Service

Write a service as you would any other [go-micro](https://github.com/micro/go-micro) service.

```
import (
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	service := k8s.NewService(
		micro.Name("my.service")
	)
	service.Init()
	service.Run()
}
```

### Install Micro

```
go get github.com/micro/kubernetes/cmd/micro
```

or

```
docker pull microhq/micro:kubernetes
```

### Healthchecking Sidecar

The healthchecking sidecar exposes `/health` as a http endpoint and calls the rpc endpoint `Debug.Health` on a service


```
go get github.com/micro/health
```

or

```
docker pull microhq/health:kubernetes
```

Run e.g healthcheck greeter service with address localhost:9091

```
health --server_name=greeter --server_address=localhost:9091
```

Call the healthchecker on localhost:8080

```
curl http://localhost:8080/health
```

Add to a kubernetes deployment

```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  namespace: default
  name: greeter
spec:
  replicas: 1
  template:
    metadata:
      labels:
        name: greeter-srv
        micro: go.micro.srv.greeter
    spec:
      containers:
        - name: greeter
          command: [
		"/greeter-srv",
		"--server_address=0.0.0.0:9091",
		"--broker_address=0.0.0.0:10001"
	  ]
          image: microhq/greeter-srv:kubernetes
          imagePullPolicy: Always
          ports:
          - containerPort: 9091
            name: greeter-port
        - name: health
          command: [
		"/health",
                "--health_address=0.0.0.0:8080",
		"--server_name=greeter",
		"--server_address=0.0.0.0:9091"
	  ]
          image: microhq/health:kubernetes
          livenessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 3
            periodSeconds: 3
```
