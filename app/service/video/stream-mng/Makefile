SHELL=/usr/bin/env bash

git=$(subst -,.,$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null || git rev-parse --short HEAD 2>/dev/null))


ldflags=-X=go-common/app/service/video/stream-mng/core.CurrentCommit='+$(git)'
ifneq ($(strip $(LDFLAGS)),)
	ldflags+=-extldflags=$(LDFLAGS)
endif
GOFLAGS+=-ldflags="$(ldflags)"

all: clean stream-mng


show-env:
	@echo '_________________build_environment_______________'
	@echo '| git commit=$(git)'
	@echo '-------------------------------------------------'

stream-mng:show-env $(BUILD_DEPS)
	go build $(GOFLAGS) -o stream-mng ./cmd/*.go

lint:
	gofmt -s -w ./
	golangci-lint run

linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o stream-mng ./cmd/*.go

clean:
	rm -rf stream-mng

.PHONY: clean
