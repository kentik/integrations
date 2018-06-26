export GOPATH := $(GOPATH):$(PWD)/vendor

.PHONY: all
all:
	bin/version.sh
	go install ./cmd/... ./pkg/...

.PHONY: test
test:
	go test ./cmd/... ./pkg/...
