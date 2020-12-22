GOPKG ?=	moul.io/progress
DOCKER_IMAGE ?=	moul/progress
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ progress hello world' > .tmp/usage.txt
	progress hello world 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
