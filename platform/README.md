# Platform folder
All the configuration required to manage a development platform (for personal usage). 
Ruling via [Makefile](./Makefile) 

## Kubernetes
Kind clusters

**Targets**

* kubernetes-create: creates `kind` cluster using [cluster.yaml](./cluster.yaml)
* kubernetes-delete: deletes `kind` cluster

## Fluxcd

* flux-install: installs fluxcd in a running kube clusters (via kubeconfig) 



