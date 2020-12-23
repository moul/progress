GOPKG ?=	moul.io/progress
DOCKER_IMAGE ?=	moul/progress
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	embedmd -w README.md
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
