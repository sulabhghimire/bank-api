# GO Bank API

Bank API is a Go-based web API designed to simulate a simple banking system. It utilizes modern tools and frameworks to manage database migrations, API development, and SQL query generation.

---

## Features

- **Database Migrations**: Managed using [Goose](https://github.com/pressly/goose) for version control and schema updates.
- **Web Framework**: Uses [Fiber](https://github.com/gofiber/fiber) for high-performance API development.
- **SQL Query Generation**: Leverages [SQLC](https://sqlc.dev/) to generate type-safe database queries from SQL files.
- **GO Mock**: Leverages [GO MOCK](https://github.com/golang/mock) to mock database.
- **Makefile Support**: Automates tasks like migrations, building, and running the application.

---

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/) (1.20 or later)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Goose](https://github.com/pressly/goose) (`go install github.com/pressly/goose/v3/cmd/goose@latest`)
- [SQLC](https://sqlc.dev/) (`go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`)
- [GO MOCK](https://github.com/golang/mock) (`go install github.com/golang/mock/mockgen@v1.6.0`)

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/your-username/bank-api.git
cd bank-api
```

### Initialize the Go Module

```bash
go mod init github.com/your-username/bank-api
go mod tidy
```

### Setup Docker Services

```bash
make docker-up
make database-create
make migrate
```
