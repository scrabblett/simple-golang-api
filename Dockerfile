FROM golang:latest
LABEL authors="Dmitry Sagan"
LABEL description="Simple api written in Go. For any questions: telegram @quryy"

WORKDIR /app
COPY . .

# Install linter
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2

# Run makefile
RUN make all

# Run backend
CMD ["go run ./cmd/app/main.go"]