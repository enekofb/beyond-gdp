kubernetes-create:
	kind create cluster --config cluster.yaml

kubernetes-delete:
	kind delete cluster

flux-install:
	flux bootstrap github \
      --owner=$(GITHUB_USERNAME) \
      --repository=desired-state \
      --branch=main \
      --path=./clusters/kind \
      --personal