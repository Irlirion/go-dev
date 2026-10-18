GO ?= go
TASK_DIRS := $(if $(strip $(DIR)),$(DIR),$(patsubst %/go.mod,%,$(wildcard hm2/*/go.mod)))

.PHONY: help test fmt vet build clean

help:
	@printf '%s\n' 'Available targets (set DIR=path to run in one Go module):' \
	  '  make test   - run tests in selected/all modules' \
	  '  make fmt    - format Go files in selected/all modules' \
	  '  make vet    - run go vet in selected/all modules' \
	  '  make build  - build selected/all modules' \
	  '  make clean  - remove local Go build and test outputs'

test:
	@set -e; for dir in $(TASK_DIRS); do \
	  test -f "$$dir/go.mod" || { printf 'No go.mod found in %s\n' "$$dir" >&2; exit 2; }; \
	  printf '\n==> %s: go test ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) test ./...); \
	done

fmt:
	@set -e; for dir in $(TASK_DIRS); do \
	  test -f "$$dir/go.mod" || { printf 'No go.mod found in %s\n' "$$dir" >&2; exit 2; }; \
	  printf '\n==> %s: go fmt ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) fmt ./...); \
	done

vet:
	@set -e; for dir in $(TASK_DIRS); do \
	  test -f "$$dir/go.mod" || { printf 'No go.mod found in %s\n' "$$dir" >&2; exit 2; }; \
	  printf '\n==> %s: go vet ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) vet ./...); \
	done

build:
	@set -e; for dir in $(TASK_DIRS); do \
	  test -f "$$dir/go.mod" || { printf 'No go.mod found in %s\n' "$$dir" >&2; exit 2; }; \
	  printf '\n==> %s: go build ./...\n' "$$dir"; \
	  (cd "$$dir" && $(GO) build ./...); \
	done

clean:
	@set -e; for dir in $(TASK_DIRS); do \
	  test -f "$$dir/go.mod" || { printf 'No go.mod found in %s\n' "$$dir" >&2; exit 2; }; \
	  printf '\n==> cleaning %s\n' "$$dir"; \
	  (cd "$$dir" && $(GO) clean && rm -f *.test coverage.out); \
	done
