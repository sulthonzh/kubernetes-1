# Micro on Kubernetes

This repo provides the config to run the micro platform on kubernetes

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
go.micro.srv.event
go.micro.srv.geo
go.micro.srv.monitor
go.micro.srv.router
go.micro.srv.trace
go.micro.web
go.micro.web.discovery
go.micro.web.event
go.micro.web.geo
go.micro.web.monitor
go.micro.web.trace
topic:geo.location
topic:micro.config.watch
topic:micro.discovery.heartbeat
topic:micro.discovery.watch
topic:micro.event.record
topic:micro.monitor.healthcheck
topic:micro.monitor.stats
topic:micro.monitor.status
topic:micro.trace.span
topic:platform.router.stats
```

### Screenshots
-
![1](https://github.com/micro/kubernetes/blob/master/doc/1.png)
-
![2](https://github.com/micro/kubernetes/blob/master/doc/2.png)
-
![3](https://github.com/micro/kubernetes/blob/master/doc/3.png)
-
![4](https://github.com/micro/kubernetes/blob/master/doc/4.png)

