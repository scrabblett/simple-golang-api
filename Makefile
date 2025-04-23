all: deps vet fmt tools deps mocks test test_coverage

deps:
	@echo Install dependencies
	go mod tidy
	go mod download

vet:
	go vet ./...

fmt:
	@echo "$(OK_COLOR)Check fmt$(NO_COLOR)"
	@echo "FIXME go fmt does not format imports, should be fixed"
	@go fmt ./...

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/vektra/mockery/v3@v3.2.3
	go get golang.org/x/tools/cmd/cover

mocks:
	go generate ./...

test:
	go test ./... -coverprofile cover.out

test_coverage:
	 go tool cover -html cover.out
