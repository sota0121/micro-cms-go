UNAME := $(shell uname)
XARGS = xargs
ARCH ?= $(shell go env GOARCH)

ifeq ($(UNAME), Linux)
XARGS = xargs -r
endif
XARGS += rm -r

.PHONY: all

build:
	GO_BUILD_FLAGS="-v" ./scripts/build.sh

clean:
	rm -rf ./bin
