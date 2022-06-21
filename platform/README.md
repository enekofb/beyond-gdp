# Platform folder
All the configuration required to manage a development platform (for personal usage). 
Ruling via [Makefile](./Makefile) 

## Kubernetes
Kind clusters

**Targets**

* kubernetes-create-dev: creates `kind` cluster using [cluster.yaml](clusters/dev.yaml)
* kubernetes-delete-dev: deletes `kind` cluster
* kubernetes-create-test: create test eks cluster using [cluster.yaml](clusters/test.yaml)
* kubernetes-delete-test: delete test eks cluster

## Fluxcd

* flux-install: installs fluxcd in a running kube clusters (via kubeconfig) 

## DNS

**beyondgdp.co.uk**

- domain registered with Ionos
- subdomain for application
  - worldhappiness.beyondgdp.co.uk



