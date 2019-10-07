TEST?=./...
PKG_NAME=slack
GO_FILES?=$$(find . -name "*.go" | grep -v vendor)
build:
	go build -v

install: build
	mkdir -p ~/.terraform.d/plugins/
	mv terraform-provider-slack ~/.terraform.d/plugins/

test:
	go test $(TEST) -timeout=30s -parallel=4 -v

testacc:
	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 go test $(TEST) -v $(TESTARGS) -timeout 240m -ldflags="-X=github.com/terraform-providers/terraform-provider-slack/version.ProviderVersion=acc"

fmt:
	gofmt -s -w $(GO_FILES)

fmtcheck:
	@sh "$(PWD)/scripts/gofmtcheck.sh"
tools:
	GO111MODULE=on go install github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=on go install github.com/client9/misspell/cmd/misspell
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint

vendor:
	GO111MODULE=on go mod vendor
.PHONY: build install test testacc fmt fmtcheck tools vendor