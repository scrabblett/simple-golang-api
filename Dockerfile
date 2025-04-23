FROM golang:latest
LABEL authors="Dmitry Sagan"
LABEL description="Simple api written in Go. For any questions: telegram @quryy"

WORKDIR /app
COPY . .

# Run makefile
RUN make all

# Run backend
CMD ["go run ./cmd/app/main.go"]