export GOPATH := $(PWD)

.PHONY: all
all:
	bin/version.sh
	go install cmd/gcevpc

#.PHONY: generate
#generate:
#	$(GOCMD) generate chf/...

.PHONY: clean
clean:
	rm -f bin/gcevpc

.PHONY: test
test:
	$(GOCMD) test chf/...
