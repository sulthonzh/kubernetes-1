# Micro on Kubernetes [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/kubernetes/go/micro?status.svg)](https://godoc.org/github.com/micro/kubernetes) [![Travis CI](https://api.travis-ci.org/micro/kubernetes.svg?branch=master)](https://travis-ci.org/micro/kubernetes) [![Go Report Card](https://goreportcard.com/badge/micro/kubernetes)](https://goreportcard.com/report/github.com/micro/kubernetes)

Micro on Kubernetes is kubernetes native micro.

Micro is a microservice toolkit. Kubernetes is a container orchestrator.

Together they provide the foundations for a microservice platform.

## Features

- No external dependencies
- Client side discovery caching
- Optional k8s service load balancing
- gRPC transport protocol
- Pre-initialised toolkit

## Getting Started

- [Installing Micro](#installing-micro)
- [Writing a Service](#writing-a-service)
- [Deploying a Service](#deploying-a-service)
- [Healthchecking Sidecar](#healthchecking-sidecar)
- [K8s Load Balancing](#k8s-load-balancing)
- [K8s Load Balancing using service-mesh](#integrating-with-conduit-service-mesh)
- [Contribute](#contribute)

## Installing Micro


```
go get github.com/micro/kubernetes/cmd/micro
```

or

```
docker pull microhq/micro:kubernetes
```

For go-micro

```
import "github.com/micro/kubernetes/go/micro"
```

## Writing a Service

Write a service as you would any other [go-micro](https://github.com/micro/go-micro) service.

```go
import (
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	service := k8s.NewService(
		micro.Name("greeter")
	)
	service.Init()
	service.Run()
}
```

## Deploying a Service

Here's an example k8s deployment for a micro service

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
        app: greeter-srv
    spec:
      containers:
        - name: greeter
          command: [
		"/greeter-srv",
		"--server_address=0.0.0.0:8080",
		"--broker_address=0.0.0.0:10001"
	  ]
          image: microhq/greeter-srv:kubernetes
          imagePullPolicy: Always
          ports:
          - containerPort: 8080
            name: greeter-port
```

Deploy with kubectl

```
kubectl create -f greeter.yaml
```

## Healthchecking Sidecar

The healthchecking sidecar exposes `/health` as a http endpoint and calls the rpc endpoint `Debug.Health` on a service. 
Every go-micro service has a built in Debug.Health endpoint.

### Install healthchecker

```
go get github.com/micro/health
```

or

```
docker pull microhq/health:kubernetes
```

### Run healtchecker

Run e.g healthcheck greeter service with address localhost:9091

```
health --server_name=greeter --server_address=localhost:9091
```

Call the healthchecker on localhost:8080

```
curl http://localhost:8080/health
```

### K8s Deployment

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
        app: greeter-srv
    spec:
      containers:
        - name: greeter
          command: [
		"/greeter-srv",
		"--server_address=0.0.0.0:8080",
		"--broker_address=0.0.0.0:10001"
	  ]
          image: microhq/greeter-srv:kubernetes
          imagePullPolicy: Always
          ports:
          - containerPort: 8080
            name: greeter-port
        - name: health
          command: [
		"/health",
                "--health_address=0.0.0.0:8081",
		"--server_name=greeter",
		"--server_address=0.0.0.0:8080"
	  ]
          image: microhq/health:kubernetes
          livenessProbe:
            httpGet:
              path: /health
              port: 8081
            initialDelaySeconds: 3
            periodSeconds: 3
```

## K8s Load Balancing

Micro includes client side load balancing by default but kubernetes also provides Service load balancing strategies. 
We can offload load balancing to k8s by using the [static selector](https://github.com/micro/go-plugins/tree/master/selector/static) 
and k8s services.

Rather than doing address resolution, the static selector returns the service name plus a fixed port e.g greeter returns greeter:8080. 
Read about the [static selector](https://github.com/micro/go-plugins/tree/master/selector/static).

This approach handles both initial connection load balancing and health checks since Kubernetes services stop routing traffic to unhealthy services, but if you want to use long lived connections such as the ones in gRPC protocol, a service-mesh like [Conduit](https://conduit.io/), [Istio](https://istio.io) and [Linkerd](https://linkerd.io/) should be prefered to handle service discovery, routing and failure. 

### Usage

To use the static selector when running your service specify the flag or env var 

```
MICRO_SELECTOR=static ./service
```

or

```
./service --selector=static
```

### K8s Deployment

An example deployment

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
        app: greeter-srv
    spec:
      containers:
        - name: greeter
          command: [
		"/greeter-srv",
		"--selector=static",
		"--server_address=0.0.0.0:8080",
		"--broker_address=0.0.0.0:10001"
	  ]
          image: microhq/greeter-srv:kubernetes
          imagePullPolicy: Always
          ports:
          - containerPort: 8080
            name: greeter-port
```

Deploy with kubectl

```
kubectl create -f deployment-static-selector.yaml
```

### K8s Service

The static selector offloads load balancing to k8s services. So ensure you create a k8s Service for each micro service. 

Here's a sample service

```
apiVersion: v1
kind: Service
metadata:
  name: greeter
  labels:
    app: greeter
spec:
  ports:
  - port: 8080
    protocol: TCP
  selector:
    app: greeter
```

Deploy with kubectl

```
kubectl create -f service.yaml
```

Calling micro service "greeter" from your service will route to the k8s service greeter:8080.

## Integrating with Conduit service-mesh
**ATTENTION**: Conduit is under heavy development and is not currently production ready.

In order to install conduit in your cluster you should first install Conduit CLI using

```curl https://run.conduit.io/install | sh```

And finnaly add Conduit CLI binary to your $PATH.

```export PATH=$PATH:$HOME/.conduit/bin```

To install conduit you need a kubernetes cluster running version 1.8 or later. To setup RBAC clusterroles for conduit-controller, web dashboard, prometheus and grafana deployments run

```conduit install | kubectl apply -f -```

To check for conduit status run

```conduit check```

Once every component have been started you are able to start running services using conduit service mesh. To access conduit web dashboard where you can see your service mesh run 

```conduit dashboard```

To start deploying apps to use conduit it is important to use [static selector](https://github.com/micro/go-plugins/tree/master/selector/static) because conduit and other service meshes use kubernetes services as a service discovery mechanism.

To deploy greeter service with health checking and conduit sidecar you will not need to change anything. As you see this deployment script is similar to last one.

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
        app: greeter-srv
    spec:
      containers:
        - name: greeter
          command: [
		"/greeter-srv",
	  "--selector=static",
		"--server_address=0.0.0.0:8080",
		"--broker_address=0.0.0.0:10001"
	  ]
          image: microhq/greeter-srv:kubernetes
          imagePullPolicy: Always
          ports:
          - containerPort: 8080
            name: greeter-port
        - name: health
          command: [
		"/health",
                "--health_address=0.0.0.0:8081",
		"--server_name=greeter",
		"--server_address=0.0.0.0:8080"
	  ]
          image: microhq/health:kubernetes
          livenessProbe:
            httpGet:
              path: /health
              port: 8081
            initialDelaySeconds: 3
            periodSeconds: 3
```

Use conduit inject to inject conduit-init container that will setup conduit-proxy's sidecar. Kubernetes will start to proxy traffic throught conduit-proxy that will handle discovery, visibility, failures..

```
conduit inject deployment.yaml | kubectl apply -f -
```

Now lets create a sample kubernetes service

```
apiVersion: v1
kind: Service
metadata:
  name: greeter
  labels:
    app: greeter
spec:
  ports:
  - port: 8080
    protocol: TCP
  selector:
    app: greeter
```

Deploy with kubectl

```
kubectl create -f service.yaml
```

Now your deployment is completed. Go to conduit's dashboard to look for this deployment and to check for inbound and outbound connections.

*If your service uses Websockets, MySQL and other protocols please read [conduit docs](https://conduit.io/adding-your-service/).*


## Contribute

We're looking for contributions from the community to help guide the development of Micro on Kubernetes

### TODO

- Integrate metaparticle - the ability to create self deploying microservices
- Example multi-service app - provide a microservice app example
- K8s api extenstons - provide config to setup the micro api as a k8s extension
- Support for micro functions
- Easy deployment scripts
