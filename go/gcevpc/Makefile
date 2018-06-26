export GOPATH := $(GOPATH):$(PWD)/vendor

.PHONY: all
all:
	go install ./cmd/... ./pkg/...

.PHONY: test
test:
	go test ./cmd/... ./pkg/...
