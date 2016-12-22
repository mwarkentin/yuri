help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Format codebase
	gofmt -w src/*

install: ## Install yuri
	go install src/yuri.go

.PHONY: help fmt install
.DEFAULT_GOAL := help
