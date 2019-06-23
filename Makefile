CMDS := $(shell go list ./cmd/...)
PKGS := $(shell go list ./pkg/...)
BIN_DIR := bin
LINTER_URI := https://install.goreleaser.com/github.com/golangci/golangci-lint.sh

.PHONY: build lint test clean $(CMDS)

build: $(CMDS)

$(CMDS):
	go build -o $(BIN_DIR)/$(shell basename $@) $@

lint:
	golangci-lint run

test: lint
	go test -timeout 1s -v $(CMDS) $(PKGS) 2>&1 | go-junit-report

clean:
	rm -rf $(BIN_DIR)

deps:
	go get ./...
	go get github.com/jstemmer/go-junit-report
	go get github.com/t-yuki/gocover-cobertura
	curl -sfL $(LINTER_URI)  | sh -s -- -b $(shell go env GOPATH)/bin v1.17.1
