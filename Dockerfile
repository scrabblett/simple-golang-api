FROM golang:latest
LABEL authors="Dmitry Sagan"

WORKDIR /app
COPY . .

RUN go build ./cmd/app/main.go

CMD ["go run ./cmd/app/main.go"]