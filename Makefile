# Borrowed from https://github.com/technosophos/helm-template/blob/master/Makefile
HELM_HOME ?= $(shell helm home)
HELM_PLUGIN_DIR ?= $(HELM_HOME)/plugins/helm-commandeer
HAS_DEP := $(shell command -v dep;)
VERSION := $(shell sed -n -e 's/version:[ "]*\([^"]*\).*/\1/p' plugin.yaml)
DIST := $(CURDIR)/_dist
LDFLAGS := "-X main.version=${VERSION}"

.PHONY: install
install: bootstrap build
	cp commandeer $(HELM_PLUGIN_DIR)
	cp plugin.yaml $(HELM_PLUGIN_DIR)

.PHONY: build
build:
	go build -o commandeer -ldflags $(LDFLAGS) ./main.go

.PHONY: dist
dist: bootstrap
	mkdir -p $(DIST)
	@echo "Building binaries in parallel"
	bash -c '\
		GOOS=linux GOARCH=amd64 go build -o commandeer -ldflags $(LDFLAGS) ./main.go && tar -zcvf $(DIST)/helm-commandeer-linux-$(VERSION).tgz commandeer README.md LICENSE plugin.yaml & \
		GOOS=darwin GOARCH=amd64 go build -o commandeer -ldflags $(LDFLAGS) ./main.go && tar -zcvf $(DIST)/helm-commandeer-macos-$(VERSION).tgz commandeer README.md LICENSE plugin.yaml & \
		GOOS=windows GOARCH=amd64 go build -o commandeer.exe -ldflags $(LDFLAGS) ./main.go && tar -zcvf $(DIST)/helm-commandeer-windows-$(VERSION).tgz commandeer.exe README.md LICENSE plugin.yaml & \
		wait \
	'

.PHONY: bootstrap
bootstrap:
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure