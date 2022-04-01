GOFILES=$(shell find . -name "*.go" -type f)
EXAMPLE_HELLOWORLD_DIR=examples/helloworld
PROTO_FILE=$(EXAMPLE_HELLOWORLD_DIR)/api/helloworld.proto
3RD_DIR=$(shell echo $$GOPATH | cut -d':' -f1)/src/go-common/3rd

protoc-gen-bm: $(GOFILES)
	go install .

protoc-gen-go: $(PROTO_FILE)
	protoc -I. -I$(3RD_DIR) --go_out=plugins=grpc:. $<

examples-helloworld: $(PROTO_FILE) protoc-gen-bm protoc-gen-go
	protoc -I. -I$(3RD_DIR) --bm_out=logtostderr=1,v=10:. $<

examples-helloworld-jsonpb: $(PROTO_FILE) protoc-gen-bm protoc-gen-go
	protoc -I. -I$(3RD_DIR) --bm_out=logtostderr=1,v=10,jsonpb=true:. $<

test:
	go test ./...

.PHONY: protoc-gen-bm protoc-gen-go examples-helloworld test
