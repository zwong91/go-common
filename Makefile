SHELL=/usr/bin/env bash

git=$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))


ldflags=-X=bucky.com/bili/core.CurrentCommit='+$(git)'
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif
GOFLAGS+=-ldflags="$(ldflags)"

all: clean discovery


show-env:
	@echo '_________________build_environment_______________'
	@echo '| git commit=$(git)'
	@echo '-------------------------------------------------'

discovery:show-env $(BUILD_DEPS)
	go build $(GOFLAGS) -o discovery ./app/infra/discovery/cmd/*.go

lint:
	gofmt -s -w ./
	golangci-lint run

linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o discovery ./app/infra/discovery/cmd/*.go

clean:
	rm -rf discovery

.PHONY: clean
