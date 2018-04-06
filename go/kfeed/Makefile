export GOPATH := $(PWD)

GOCMD := go

EXECUTABLES := kfeed

.PHONY: all
all:
	./bin/version.sh
	$(GOCMD) install $(addprefix chf/cmd/,$(EXECUTABLES))

.PHONY: generate
generate:
	$(GOCMD) generate chf/...

make-%:
	./bin/version.sh
	$(GOCMD) install chf/cmd/$*

.PHONY: clean
clean:
	cd bin \
		&& rm -f $(EXECUTABLES)
	rm -rf pkg
	rm -f ./bin/debs/*

#---------------------------------------------------------

.PHONY: test
test:
	$(GOCMD) test chf/...
