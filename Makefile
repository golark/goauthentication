.PHONY: dockerbuild, dockerrun, dockerstop, run, proto, clean, lint, genkey, doc, docs

SHELL := /bin/bash
DOCKER_SERVICE=goauthenticate
APP_NAME=goauthenticate
WEB_APP=$(APP_NAME)_web
API_APP=$(APP_NAME)_api
API_URL='http://localhost:5023'

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
	@echo "starting the apps"
	@source env.sh && ./dist/$(API_APP) &
	@source env.sh && ./dist/$(WEB_APP) -apiurl=$(API_URL)

test:
	@go test ./app -v
	@go test ./web -v

build: clean buildapi buildweb
	@echo "built binaries"

clean:
	@echo "clean binaries under ./dist"
	@- rm -f dist/*
	@go clean

buildweb:
	@go build -o dist/$(WEB_APP) ./web
	@echo "built $(WEB_APP)"

buildapi:
	@go build -o dist/$(API_APP) ./api
	@echo  "built $(API_APP)"

stop:
	@-pkill SIGTERM -f
