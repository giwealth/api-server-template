SHELL:=/bin/sh

export GO111MODULE=on
export GOPROXY=https://goproxy.haochang.tv,https://goproxy.cn,direct

# Path Related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
DIST_DIR := ${MKFILE_DIR}artifacts
ProjectUrl=$(shell git remote get-url origin | sed -e "s/.\{4\}//" -e "s/\.git//" -e "s/:/\//")
ProjectName=$(shell basename ${ProjectUrl})
ProjectNameDefault=$(if ${ProjectName},${ProjectName},"app" )
BuildAuthor=$(shell git config user.email)
BuildBPipeline=${if ${CI_PIPELINE_URL},${CI_PIPELINE_URL},"null"}
BuildBranch=$(shell git branch --show-current)
BuildBranchDefault=${if ${BuildBranch},${BuildBranch},"default-branch"}
BuildTime=$(shell date "+%Y-%m-%d %H:%M:%S")
CommitID=$(shell git rev-parse --short HEAD)
CommitIDDefault=${if ${CommitID},${CommitID},"default-commitID"}
ArtifactDir=${DIST_DIR}/${ProjectNameDefault}
GO_LD_FLAGS=-ldflags "\
	-X 'main.commit=${CommitIDDefault}' \
	-X 'main.branch=${BuildBranchDefault}' \
	-X 'main.buildTime=${BuildTime}' \
	-X 'main.buildAuthor=${BuildAuthor}' \
	-X 'main.pipeline=${BuildBPipeline}'"


## build: build server binary
.PHONY: build
build: revive wire
	@echo "build linux server"
	cd ${MKFILE_DIR} && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && \
	go build -v -trimpath ${GO_LD_FLAGS} -o ${DIST_DIR}/${ProjectNameDefault} ${MKFILE_DIR}cmd/server/main.go

## test: run unit tests
.PHONY: test
test: revive wire
	cd ${MKFILE_DIR}
	go test -v ./... | grep -v '^?'

## serve: start server
.PHONY: serve
serve: revive wire
	go run cmd/server/main.go --config configs/config.dev.yaml api

## mod_update: update go modules
.PHONY: mod_update
mod_update:
	cd ${MKFILE_DIR} && go get -u -v ./...


## wire: generates wire_gen.go
.PHONY: wire
wire:
	wire gen ./internal/...

## install: install dependencies
install:
	@go install github.com/google/wire/cmd/wire@latest

## revive: checks for errors in Go source code
.PHONY: revive
revive:
	revive -config revive.toml -exclude ./vendor/... ./...

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

