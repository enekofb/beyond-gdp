kubernetes-create-dev:
	kind create cluster --config clusters/dev.yaml

kubernetes-delete-dev:
	kind delete cluster

kubernetes-create-test:
	eksctl create cluster -f clusters/test.yaml

kubernetes-delete-test:
	eksctl delete cluster -f clusters/test.yaml

kubeconfig-test:
	aws eks update-kubeconfig --region eu-west-1 --name beyondgdp

flux-install:
	flux bootstrap github \
      --owner=$(GITHUB_USERNAME) \
      --repository=desired-state \
      --branch=main \
      --path=./clusters/kind \
      --personal