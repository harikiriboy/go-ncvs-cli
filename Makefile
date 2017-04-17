.PHONY: \
	build \
	clean \
	dep \
	fmt \
	fmtcheck \
	install \
	lint \
	pretest \
	test \
	vet \

all: dep build test

SRCS = $(shell git ls-files '*.go')
PKGS = ./. ./commands ./ncvs ./writer

VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

build: main.go
	go build -ldflags "$(LDFLAGS)" -o ncvs-cli $<

clean:
	$(foreach pkg,$(PKGS),go clean $(pkg) || exit;)

dep:
	@go get -u github.com/golang/dep/...
	dep ensure

fmt:
	gofmt -w $(SRCS)

fmtcheck:
	$(foreach file,$(SRCS),gofmt -d $(file);)

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

pretest: lint vet fmtcheck

test: pretest
	$(foreach pkg,$(PKGS),go test -v $(pkg) || exit;)

vet:
	$(foreach pkg,$(PKGS),go vet $(pkg);)
