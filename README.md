# Micro on Kubernetes

This repo provides the config to run the Micro on Kubernetes

An example deployment is currently running at [web.micro.pm](http://web.micro.pm) 
leveraging Google Container Engine.

Note: This is by no means a highly available deployment. It's for testing purposes. 
In the future we'll lay out a highly scalable and fault tolerant regional and 
global deployment model.

## What's in the repo?

Currently config to run Micro on Kubernetes (Testing on Google Container Engine).

- Registry - Service Discovery using consul
- Database - Percona with Galera (single cluster for now)
- Micro - API, Web UI and Sidecar (spins up GCE Load Balancers)
- Platform - All the micro platform services and dashboards
- Misc - other miscellaneous services

## Demo

Micro Demo (SSL soon):

- API - [api.micro.pm](http://api.micro.pm)
- Web UI - [web.micro.pm](http://web.micro.pm)
- Sidecar - [proxy.micro.pm](http://proxy.micro.pm)

### Play via CLI

The Micro CLI now supports proxying via the sidecar. Try it out.

```shell
micro --proxy_address=proxy.micro.pm list services
consul
go.micro.api
go.micro.api.geo
go.micro.sidecar
go.micro.srv.auth
go.micro.srv.config
go.micro.srv.db
go.micro.srv.discovery
...
```

### Getting Started

If you want to run the demo yourself, look no further.

Here's the steps I took to get started.

###### Kubernetes on Google Container Engine

Google Container Engine is the easiest way to run a managed kubernetes cluster. What's even better? $300 free credit for 60 days.

1. Get yourself a Free Trial of Google Container Engine https://cloud.google.com/free-trial/
2. Spin up a 3-4 node Kubernetes cluster
3. Setup your kubectl command https://cloud.google.com/container-engine/docs/before-you-begin

###### Launch all the things

This repo contains all the kubernetes config files with docker images prebaked. Just kick them off and all should be well. 

If you don't want to run everything you can just start what's in the registry/ and micro/ directories. We'll start there anyway.

First you may want to change the hostnames for CORS support in the API and WEB proxy. We're running a live demo on these hostnames 
hence them being present.

```
$ find . -name *.yaml | xargs grep micro.pm
./micro/micro-api.yaml:            "--api_cors=http://api.micro.pm",
./micro/micro-web.yaml:            "--web_cors=http://web.micro.pm",
```

Ok next.

1. Start Registry
2. Start Micro

Optional:

3. Start Database
4. Start Platform
5. Start Miscellaneous Services

###### Starting things

We'll start the registry and micro first
```
for dir in registry micro; do
	pushd $dir
	for config in *.yaml; do
		kubectl create -f $config
	done
	popd
done
```

###### Start optional things

Look in database/ to see how to bootstrap the DB. It requires running some commands for databases, tables and grants. 
Otherwise use same method as above to start the database, platform and other services.

### Screenshots
![1](https://github.com/micro/kubernetes/blob/master/doc/1.png)
-
![2](https://github.com/micro/kubernetes/blob/master/doc/2.png)
-
![3](https://github.com/micro/kubernetes/blob/master/doc/3.png)
-
![4](https://github.com/micro/kubernetes/blob/master/doc/4.png)

