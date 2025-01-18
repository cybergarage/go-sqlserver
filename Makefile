# Copyright (C) 2024 The go-sqlserver Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

PREFIX?=$(shell pwd)

GOBIN := $(shell go env GOPATH)/bin
PATH := $(GOBIN):$(PATH)

GIT_ROOT=github.com/cybergarage
PRODUCT_NAME=go-sqlserver
MODULE_ROOT=${GIT_ROOT}/${PRODUCT_NAME}

PKG_NAME=sql
PKG_VER=$(shell git describe --abbrev=0 --tags)
PKG_COVER=${PKG_NAME}server-cover
PKG_SRC_ROOT=${PKG_NAME}
PKG=${MODULE_ROOT}/${PKG_SRC_ROOT}

TEST_SRC_ROOT=${PKG_NAME}test
TEST_PKG=${MODULE_ROOT}/${TEST_SRC_ROOT}

BINS_ROOT=cmd
BINS_SRC_ROOT=${BINS_ROOT}
BINS_DEAMON_BIN=${PRODUCT_NAME}
BINS_PKG_ROOT=${GIT_ROOT}/${PRODUCT_NAME}/${BINS_ROOT}
EXAMPLE_BINARIES=\
	${BINS_PKG_ROOT}/${BINS_DEAMON_BIN}

EXAMPLE_DOCKER_TAG=cybergarage/${BINS_DEAMON_BIN}:${PKG_VER}
EXAMPLE_DOCKER_TAG_LATEST=cybergarage/${BINS_DEAMON_BIN}:latest

BINARIES=\
	${EXAMPLE_BINARIES}

.PHONY: clean test version
.IGNORE: lint

all: test

version:
	@pushd ${PKG_SRC_ROOT} && ./version.gen > version.go && popd
	-git commit ${PKG_SRC_ROOT}/version.go -m "Update version"

format: version
	gofmt -s -w ${PKG_SRC_ROOT} ${TEST_SRC_ROOT} ${BINS_SRC_ROOT}

vet: format
	go vet ${PKG}

lint: vet
	golangci-lint run ${PKG_SRC_ROOT}/... ${TEST_SRC_ROOT}/... ${BINS_SRC_ROOT}/...

test: lint
	chmod og-rwx  ${TEST_SRC_ROOT}/certs/key.pem
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

test_only:
	chmod og-rwx  ${TEST_SRC_ROOT}/certs/key.pem
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

build: vet
	go build -v ${BINARIES}

install:
	go install -v -gcflags=${GCFLAGS} ${BINARIES}

run: install
	${GOBIN}/${BINS_DEAMON_BIN} --debug

image: test
	docker image build -t${EXAMPLE_DOCKER_TAG_LATEST} -t${EXAMPLE_DOCKER_TAG} .
	docker push ${EXAMPLE_DOCKER_TAG_LATEST}

image-push: test
	docker push ${EXAMPLE_DOCKER_TAG}

rund: image
	docker container run -it --rm -p 5432:5432 ${EXAMPLE_DOCKER_TAG}

%.md : %.adoc
	asciidoctor -b docbook -a leveloffset=+1 -o - $< | pandoc -t markdown_strict --wrap=none -f docbook > $@
csvs := $(wildcard doc/*/*.csv)
docs := $(patsubst %.adoc,%.md,$(wildcard doc/*.adoc))
doc_touch: $(csvs)
	touch doc/*.adoc

doc: doc_touch $(docs)

clean:
	go clean -i ${PKG}
