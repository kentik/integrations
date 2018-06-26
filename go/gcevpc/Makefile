export GOPATH := $(GOPATH):$(PWD)/vendor

.PHONY: all
all:
	bin/version.sh
	go install ./cmd/... ./pkg/...

.PHONY: clean
clean:
	rm -f bin/gcevpc

.PHONY: test
test:
	$(GOCMD) test chf/...
