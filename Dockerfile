FROM golang:latest
LABEL authors="Dmitry Sagan"

WORKDIR /app
COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build ./cmd/app/main.go

CMD ["go run ./cmd/app/main.go"]