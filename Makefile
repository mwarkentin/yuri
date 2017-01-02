help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Format codebase
	gofmt -w src/*

install: ## Install yuri
	go install src/yuri.go

release: ## Publish release to equinox.io | args: version
	ifndef version
	  $(error version is undefined)
	endif
	@echo equinox release --config \"./equinox.yaml\" --version \"${version}\" --token \"****\" ./src/yuri.go
	@equinox release --config "./equinox.yaml" --version "${version}" --token "${EQUINOX_TOKEN}" ./src/yuri.go

test: ## Run tests
	go test -v -cover ./...

vendor: ## Update dependencies
	govendor fetch github.com/urfave/cli

.PHONY: help fmt install vendor
.DEFAULT_GOAL := help
