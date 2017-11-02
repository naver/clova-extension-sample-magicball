PACKAGE  = magicball
EXENAME  = magicball
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= 1
GOPATH   = $(CURDIR)/.gopath
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))

GO       = go
GOFMT    = go fmt
GOTEST   = go test -v
GLIDE    = glide
TIMEOUT  = 15

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# GIT_HEAD = $(shell git rev-parse HEAD)
# DOCKER_IMG = clova-home-hue-extension:$(GIT_HEAD)

.PHONY: all
all: fmt build

.PHONY: build
build: vendor $(BASE) ; $(info $(M) building executable to bin/ …) @
	$Q cd $(BASE) && $(GO) build \
		-i -v \
		-tags release \
		-ldflags "-X $(PACKAGE).Version=$(VERSION) -X $(PACKAGE).BuildDate=$(DATE)" \
		-o bin/$(EXENAME) cmd/main.go

$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# Tools

GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) building golint…)
	$Q go get github.com/golang/lint/golint

.PHONY: lint
lint: $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) $$($(GLIDE) novendor)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GOFMT) $$($(GLIDE) novendor)

.PHONY: test
test: $(BASE) ; $(info $(M) running go test...) @ ## Run go test
	$Q cd $(BASE) && $(GOTEST) $$($(GLIDE) novendor)

# Dependency management

glide.lock: glide.yaml | $(BASE) ; $(info $(M) updating dependencies…)
	$Q cd $(BASE) && $(GLIDE) update
	@touch $@
vendor: glide.lock | $(BASE) ; $(info $(M) retrieving dependencies…) @ ## Install Go dependencies under vendor/
	$Q cd $(BASE) && $(GLIDE) install
	@ln -sf . vendor/src
	@touch $@

# Misc

.PHONY: run
run: build run-only ; @ ## Build and run a server
	$Q bin/$(EXENAME)

.PHONY: run-only
run-only: ; $(info $(M) running a server…) @ ## Run a server.
	$Q bin/$(EXENAME)

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf glide.lock
	@rm -rf bin
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
