CACHE_MOD1 	  := $${PWD}/cache
CACHE_MOD2 	  := $${PWD}/go-cache
CACHE_MOD3 	  := $${PWD}/gocache
.DEFAULT_GOAL := help

PHONY: all help test
all: help

init: ## initialize
	@export GO111MODULE=on
	@go mod tidy

test: init ## The benchmark start.
	@sudo purge; sync; sync;
	@for CACHE_MOD in $(CACHE_MOD1) $(CACHE_MOD2) $(CACHE_MOD3); \
	do \
		echo "Benchmark target: " $$CACHE_MOD; \
		cd $$CACHE_MOD; \
    		go test -bench=. -benchmem; \
		echo "----------------------"; \
	done

help:
	@echo "Usage: make \033[36m<subCommand>\033[0m\n"
	@echo "\033[36m<subCommand>\033[0m:"
	@perl -lne ' /(.*):\s+?.*##\s+?(.*)/ and printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2' $(MAKEFILE_LIST)
