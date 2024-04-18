# Books library API

It's a simple api, representing books library. Written in Go.

## Deployement

By default, http-server starts on http://localhost:8080.

**API routes available by path /api/v1**

Example:
http://localhost:8080/api/v1/auth/login

### In Docker container

1. Install Docker
2. Run ```docker-compose up```

### Locally

1. [Install Go 1.22](https://go.dev/doc/install)
2. [Install Postgres](https://www.postgresql.org/download/) and create new database
3. Run migrations :

```bash
cd migrations/scripts
goose postgres <connection_string> up
```

Example:

```bash
goose postgres postgres://postgres:12345@localhost:5432/books_library?sslmode=disable up 
```

4. By default, configs take from env variables. [Env-config is here](.env.dev). Set variables in your system.
5. If env-variables not set, [config takes default values](internal/config/config.go)
6. Run ```go run ./cmd/app/main.go```

## Running tests

1. Run ```go test ./... -coverprofile cover.out```
2. Generate test-coverage in html (need **go cover** installed) ```go tool cover -html cover.out```

## Swagger

[Swagger is here](swagger/swagger.yaml)

1. [Go to swagger-editor](https://editor.swagger.io/)
2. Paste content of swagger.yaml in editor
