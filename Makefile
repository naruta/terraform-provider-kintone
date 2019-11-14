PLUGINS_DIR=~/.terraform.d/plugins/$(shell go env GOHOSTOS)_$(shell go env GOHOSTARCH)
SOURCES=$(wildcard *.go)
PKG_NAME=terraform-provider-kintone
# TEST?=$$(go list ./ | grep -v 'vendor')
TEST?=./...

default: clean build

build:
	go build -o $(PKG_NAME)

clean:
	rm -f $(PKG_NAME)

testacc:
	TF_ACC=1 go test $(TEST) -v

.PHONY: clean

install: build
	mkdir -p $(PLUGINS_DIR)
	cp $(PKG_NAME) $(PLUGINS_DIR)
