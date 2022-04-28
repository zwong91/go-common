SHELL=/usr/bin/env bash

all: build
.PHONY: all

unexport GOFLAGS

GOCC?=go

GOVERSION:=$(shell $(GOCC) version | tr ' ' '\n' | grep go1 | sed 's/^go//' | awk -F. '{printf "%d%03d%03d", $$1, $$2, $$3}')
ifeq ($(shell expr $(GOVERSION) \< 1016000), 1)
$(warning Your Golang version is go$(shell expr $(GOVERSION) / 1000000).$(shell expr $(GOVERSION) % 1000000 / 1000).$(shell expr $(GOVERSION) % 1000))
$(error Update Golang to version to at least 1.16.0)
endif

git=$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))

BINS:=

ldflags=-X=bucky.com/bili/core.CurrentCommit='+$(git)'
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif
GOFLAGS+=-ldflags="$(ldflags)"

show-env:
	@echo '_________________build_environment_______________'
	@echo '| git commit=$(git)'
	@echo '-------------------------------------------------'

discovery: show-env $(BUILD_DEPS)
	rm -rf discovery
	$(GOCC) build $(GOFLAGS) -o discovery ./app/infra/discovery/cmd/*.go
.PHONY: discovery
BINS+=discovery

config: show-env $(BUILD_DEPS)
	rm -rf config
	$(GOCC) build $(GOFLAGS) -o config ./app/infra/config/cmd/*.go
.PHONY: config
BINS+=config

lint:
	gofmt -s -w ./
	golangci-lint run

build: $(BINS)

clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean
