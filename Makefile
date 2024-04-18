all: mocks test test_coverage vet fmt tools lint  deps

mocks:
	go generate ./...

test:
	go test ./... -coverprofile cover.out

test_coverage:
	 go tool cover -html cover.out

vet:
	go vet ./...

fmt:
	@echo "$(OK_COLOR)Check fmt$(NO_COLOR)"
	@echo "FIXME go fmt does not format imports, should be fixed"
	@go fmt ./...

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go get golang.org/x/tools/cmd/cover

lint:
	golangci-lint run -enable-all

deps:
	@echo Install dependencies
	go mod tidy
	go mod download
