#!make
SHELL := /bin/bash
.SHELLFLAGS := -ec
AWS_REGION=eu-west-1
AWS_PROFILE=enekofb-777171359344

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

lambda-deploy:
	aws cloudformation deploy --region $(AWS_REGION) --profile $(AWS_PROFILE)  \
		--stack-name world-happiness --capabilities CAPABILITY_NAMED_IAM \
		--tags Name="world-happiness-api" Application="world-happiness-api" \
		--template-file ./lambda/cf.yaml --no-fail-on-empty-changeset
