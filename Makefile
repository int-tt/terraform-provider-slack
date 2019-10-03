TEST?=./...
PKG_NAME=slack

.PHONY: build
build:
	go build -v

.PHONY: install
install: build
	mkdir -p ~/.terraform.d/plugins/
	mv terraform-provider-slack ~/.terraform.d/plugins/

.PHONY: test
test:
	go test $(TEST) -timeout=30s -parallel=4 -v

.PHONY: fmt
fmt:
	gofmt -s -w ./$(PKG_NAME)

.PHONY: tools
tools:
	GO111MODULE=on go install github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=on go install github.com/client9/misspell/cmd/misspell
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint