#!make
SHELL := /bin/bash
.SHELLFLAGS := -ec
.DEFAULT_GOAL := build
AWS_REGION=eu-west-1
AWS_PROFILE=default

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

build: test
	go build

docker-build: build
	docker build -t world-happiness-api:latest .

docker-run:
	docker run -p 8080:8080 --name world-happiness-api --rm world-happiness-api:latest

ecr-login:
	aws ecr --region $(AWS_REGION) --profile $(AWS_PROFILE) get-login-password | docker login --username AWS --password-stdin https://436139220509.dkr.ecr.eu-west-1.amazonaws.com

docker-push: ecr-login
	docker tag world-happiness-api:latest 436139220509.dkr.ecr.eu-west-1.amazonaws.com/world-happiness-api:latest
	docker push 436139220509.dkr.ecr.eu-west-1.amazonaws.com/world-happiness-api:latest

install: test
	go install

load-image:
	kind load docker-image world-happiness-api:latest --name development

helm-upgrade:
	helm upgrade world-happiness-api .charts/world-happiness-api -n world-happiness --wait --timeout 120s --create-namespace --install

upgrade: docker-build helm-upgrade

infra-deploy:
	aws cloudformation deploy --region $(AWS_REGION) --profile $(AWS_PROFILE)  \
		--stack-name world-happiness --capabilities CAPABILITY_NAMED_IAM \
		--tags Name="world-happiness-api" Application="world-happiness-api" \
		--template-file ./.resources/cf.yaml --no-fail-on-empty-changeset

infra-dev-deploy:
	aws cloudformation deploy --region $(AWS_REGION) --profile $(AWS_PROFILE)  \
		--stack-name world-happiness-dev --capabilities CAPABILITY_NAMED_IAM \
		--tags Name="world-happiness-api-dev" Application="world-happiness-api" \
		--template-file ./.resources/dev-cf.yaml --no-fail-on-empty-changeset

export-api-spec:
	aws apigateway get-export --parameters extensions='apigateway' \
	--rest-api-id l84n6w2qs6 --stage-name prod --export-type oas30 \
	--accepts application/yaml api-spec-latest.yaml
