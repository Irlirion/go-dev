GO ?= go
TASK_DIRS := $(patsubst %/go.mod,%,$(wildcard hm2/*/go.mod))

.PHONY: help test fmt vet build clean

help:
	@printf '%s\n' 'Available targets:' \
	  '  make test   - run tests in every hm2 task' \
	  '  make fmt    - format Go files in every hm2 task' \
	  '  make vet    - run go vet in every hm2 task' \
	  '  make build  - build every hm2 task' \
	  '  make clean  - remove local Go build and test outputs'

test:
	@set -e; for dir in $(TASK_DIRS); do \
	  printf '\n==> %s: go test ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) test ./...); \
	done

fmt:
	@set -e; for dir in $(TASK_DIRS); do \
	  printf '\n==> %s: go fmt ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) fmt ./...); \
	done

vet:
	@set -e; for dir in $(TASK_DIRS); do \
	  printf '\n==> %s: go vet ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) vet ./...); \
	done

build:
	@set -e; for dir in $(TASK_DIRS); do \
	  printf '\n==> %s: go build ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) build ./...); \
	done

clean:
	@set -e; for dir in $(TASK_DIRS); do \
	  printf '\n==> cleaning %s\n' "$$dir"; \
	  (cd "$$dir" && $(GO) clean && rm -f -- *.test coverage.out); \
	done
