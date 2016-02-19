# Micro on Kubernetes

This repo provides the config to run the micro platform on kubernetes

An example deployment is currently running at [web.micro.pm](http://web.micro.pm) 
leveraging Google Container Engine.

## What's in the repo?

Currently config to run Micro on Kubernetes (Testing on Google Container Engine).

- Registry - Service Discovery using consul
- Database - Percona with Galera (single cluster for now)
- Micro - API, Web UI and Sidecar (spins up GCE Load Balancers)
- Platform - All the micro platform services
- Misc - other miscellaneous services

