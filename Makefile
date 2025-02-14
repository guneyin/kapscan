BINARY_NAME=kapscan

.PHONY: build

PACKAGE=github.com/guneyin/kapscan
VERSION=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH=$(shell git rev-list -1 HEAD)
BUILD_TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S')

LDFLAG_VERSION='${PACKAGE}/util.Version=${VERSION}'
LDFLAG_COMMIT_HASH='${PACKAGE}/util.CommitHash=${COMMIT_HASH}'
LDFLAG_BUILD_TIMESTAMP='${PACKAGE}/util.BuildTime=${BUILD_TIMESTAMP}'

init: clean tidy vet build

tidy:
	go mod tidy

vet:
	go vet ./...

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

run:
	go run .

build:
	go build -o ${BINARY_NAME} -ldflags "-X ${LDFLAG_VERSION} -X ${LDFLAG_COMMIT_HASH} -X ${LDFLAG_BUILD_TIMESTAMP}" .

clean:
	go clean
	rm -f ${BINARY_NAME}


