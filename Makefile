GIT_COMMIT := $(shell git rev-list -1 HEAD)

all: clean build

distclean: clean
clean:
	rm -f check_hp_disk_firmware

build: check_hp_disk_firmware
check_hp_disk_firmware: main.go
	go build -o $@ -ldflags "-X main.GitCommit=$(GIT_COMMIT)"
