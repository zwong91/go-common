SHELL=/usr/bin/env bash

git=$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))


ldflags=-X=go-common/app/service/ops/log-agent/core.CurrentCommit='+$(git)'
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif
GOFLAGS+=-ldflags="$(ldflags)"

all: clean log-agent


show-env:
	@echo '_________________build_environment_______________'
	@echo '| git commit=$(git)'
	@echo '-------------------------------------------------'

log-agent:show-env $(BUILD_DEPS)
	go build $(GOFLAGS) -o log-agent ./cmd/*.go

lint:
	gofmt -s -w ./
	golangci-lint run

linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o log-agent ./cmd/*.go

clean:
	rm -rf log-agent

.PHONY: clean
