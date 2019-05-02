APP?=TravisTest
USERSPACE?=MaksRybalkin/TravisTest
RELEASE?=0.0.1
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux

CONTAINER_NAME?=${APP}

REPO_INFO=$(shell git config --get remote.origin.url)
DATE = $(shell date -u +%Y%m%d.%H%M%S)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

.PHONY: all # All targets are accessible for user
.DEFAULT: help # Running Make will run the help target

help: ## Show Help
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

cover_docker: ## Running tests in docker
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

publish_cover: cover_docker ## Publish coverage to coveralls
	go get -v github.com/mattn/goveralls
	go get -v golang.org/x/tools/cmd/cover
	goveralls -v -show -coverprofile=coverage.out

dbuild: ## Build and tag containers
	docker build -t maksrybalkin/travistest:${RELEASE} -f ./Dockerfile .

dpush: dbuild  ## Push containers to docker hub
	docker push maksrybalkin/travistest:${RELEASE}

heroku_deploy: ## Build image and deploy it ot heroku
	@echo $(HEROKU_API_KEY) | docker login -u _ --password-stdin registry.heroku.com
	docker build -t registry.heroku.com/travisherokutest/web -f Dockerfile .
	docker push registry.heroku.com/travisherokutest/web
	heroku container:release web -a travisherokutest

build:
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo \
		-ldflags "-s -w -X main.RELEASE=${RELEASE} -X main.COMMIT=${COMMIT} -X main.DATE=${DATE} -X main.REPO=${REPO_INFO}" \
		-o ${APP}