.PHONY: dockerbuild, dockerrun, dockerstop, run, proto, clean, lint, genkey, doc, docs

DOCKER_SERVICE=goauthenticate
APP_NAME=goauthenticate
WEB_APP=$(APP_NAME)_web
API_APP=$(APP_NAME)_api
API_URL='http://localhost:5023'
SHELL := /bin/bash

dockerpush: dockerbuild
	docker-compose push

dockerbuild:
	docker-compose build --no-cache

dockerstop:
	docker-compose down

dockerrun:
	docker-compose run --service-ports $(DOCKER_SERVICE)

dockerinteractive:
	docker-compose run --service-ports $(DOCKER_SERVICE) sh

run: build
	source env.sh && ./dist/$(APP_NAME)_web -apiurl=$(API_URL)

test:
	@go test ./app -v
	@go test ./web -v

build: clean buildapi buildweb
	@printf "built binaries"

clean:
	@- rm -f dist/*
	@go clean

buildweb:
	@go build -o dist/$(WEB_APP) ./web

buildapi:
	@go build -o dist/$(API_NAME) ./api

stop:
	@-pkill SIGTERM -f
