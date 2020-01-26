GIT_COMMIT := $(shell git rev-list -1 HEAD)
GO_BUILD := go build -v -ldflags "-X main.GitCommit=$(GIT_COMMIT)"

.PHONY: all clean build test

all: clean build tarball

distclean: clean
clean:
	rm -rf build/

build:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o build/check_hp_disk_firmware-amd64 .
	GOOS=linux GOARCH=386 $(GO_BUILD) -o build/check_hp_disk_firmware-i386

tarball: build
	cd build && tar cf check_hp_disk_firmware.tar.gz check_hp_disk_firmware-*

test:
	go test -v ./...
