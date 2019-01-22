PLUGINS_DIR=~/.terraform.d/plugins/darwin_amd64
SOURCES=$(wildcard *.go)
PKG_NAME=terraform-provider-kintone
# TEST?=$$(go list ./ | grep -v 'vendor')
TEST?=./...

default: clean build

build:
	# dep ensure
	go build -o $(PKG_NAME)

clean:
	rm $(PKG_NAME)

testacc:
	TF_ACC=1 go test $(TEST) -v

.PHONY: clean

install: build
	mkdir -p $(PLUGINS_DIR)
	cp $(PKG_NAME) $(PLUGINS_DIR)
