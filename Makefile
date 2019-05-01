.PHONY: all # All targets are accessible for user
.DEFAULT: help # Running Make will run the help target


dtest: ## Running tests in docker
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

dbuild: ## Build and tag containers
	docker build -t maksrybalkin/travistest:0.0.1 -f ./Dockerfile .

dpush: dbuild  ## Push containers to docker hub
	docker push maksrybalkin/travistest:0.0.1

help: ## Show Help
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'