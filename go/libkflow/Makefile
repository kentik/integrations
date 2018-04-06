MAIN := github.com/kentik/libkflow
CLIB := $(MAIN)/cmd/libkflow
PKGS := $(MAIN) $(MAIN)/api $(MAIN)/chf $(CLIB)

CFLAGS  += -std=c99

OS      := $(shell go env GOOS)
ARCH    := $(shell go env GOARCH)
TARGET  := $(OS)-$(ARCH)
VERSION ?= $(shell git describe --tags --always --dirty)
WORK    := $(CURDIR)/out/$(TARGET)

ARCHIVE   := libkflow-$(VERSION)-$(TARGET).tar.gz
ARTIFACTS :=           \
    $(WORK)/libkflow.a \
    $(WORK)/server     \
    $(WORK)/demo       \
    $(MAIN)/kflow.h    \
    $(CURDIR)/demo.c

ifeq ($(OS), darwin)
	LDFLAGS += -framework Security -framework CoreFoundation
else ifeq ($(OS), linux)
	LDFLAGS += -lpthread -fPIC
endif

file-types = .GoFiles .CgoFiles .HFiles

find-files = $(foreach f,$(file-types),$(call list-files,$f,$1))
list-files = $(shell go list -f '{{range $$f := $1}}$2/{{$$f}} {{end}}' $2)

# for each package in $(PKGS) define a variable named SRC_$(pkg)
# containing all of the files in that package.
$(foreach pkg,$(PKGS),$(eval SRC_$(pkg) := $(call find-files,$(pkg))))

# define a variable named SRC containing all files in all packages.
$(foreach pkg,$(PKGS),$(eval SRC += $(SRC_$(pkg))))


$(ARCHIVE): $(ARTIFACTS)
	$(info building $(ARCHIVE))
	@tar czf $@ $(foreach f,$^,-C $(dir $f) $(notdir $f))

$(WORK)/libkflow.a: $(SRC)
	go build -o $@ -buildmode=c-archive -ldflags="-X $(MAIN).Version=$(VERSION)" $(CLIB)

$(WORK)/server: $(SRC)
	go build -o $@ $(MAIN)/cmd/server

$(WORK)/demo: $(MAIN)/kflow.h $(CURDIR)/demo.c $(WORK)/libkflow.a
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ -I $(<D) $(filter-out $<,$^)

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

clean:
	rm -f *.tar.gz
	rm -rf out/*

.PHONY: test
.PHONY: clean

.SUFFIXES:

VPATH = $(GOPATH)/src
