#!make
SHELL := /bin/bash
.SHELLFLAGS := -ec

test:
	go test ./...

build: test
	go build

docker-build: build
	docker build -t world-happiness-api:latest .

docker-run:
	docker run -p 8080:8080 --name world-happiness-api --rm world-happiness-api:latest

docker-push:
	docker push world-happiness-api:latest

integration-tests:
	go test -v ./integration/... -tags integration

install: test
	go install

helm-upgrade:
	helm upgrade world-happiness-api charts/world-happiness-api -n egctl --wait --timeout 120s --create-namespace --install

upgrade: docker-build helm-upgrade