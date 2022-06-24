.PHONY: help dbup dfup cover tidy

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

dbup: ## Docker-compose UP DB
	@docker-compose up -d postgres mysql
	@docker network inspect simples-restapi_fullstack |  jq -r '.[].Containers[] | {Name: .Name, IP: .IPv4Address}' 

dfup:  ## Docker-compose UP Front
	@docker-compose up -d app 
	@docker network inspect simples-restapi_fullstack |  jq -r '.[].Containers[] | {Name: .Name, IP: .IPv4Address}' 

tidy:
	@go clean --modcache
	@go mod tidy 

cover: ## test cover
	@echo "Running coverage"
	@go test -cover -race ./app/usecase ./app/repository

clean: ## Docker-compose Cleaning
	@docker-compose down
	docker system prune -a --volumes