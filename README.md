---
runme:
  id: 01HHMEVXQN9J74NRS7YV80H7MD
  version: v2.0
---

# CRUD in Go with GORM,Testify ,Docker Compose, Redis, and PostgreSQL

This is a simple example of a CRUD application in Go that uses GORM as the ORM, Docker Compose for container orchestration, and Redis and PostgreSQL as databases.
Configuration

Make sure you have Docker and Docker Compose installed on your machine.

Clone the repository:

```sh {"id":"01HHMEVXQKK7VAH37PACDAQP8C"}
git clone github.com/isaquerr25/crud-go-gin-redis-gorm-testify-docker

```

Create a .env file in the project root and configure the necessary environment variables. An example can be found in the .env.example file.

Start the containers using Docker Compose:

```sh {"id":"01HHMEVXQMWZ24KMH4Y70QA7NM"}
docker-compose up -d

```

Run the application:

```sh {"id":"01HHMEVXQMWZ24KMH4YAP3KPYP"}
    go run main.go

```

Routes

    GET /users: Get all users.
    GET /users/:id: Get a specific user.
    POST /users: Create a new user.
    PUT /users/:id: Update an existing user.
    DELETE /users/:id: Delete a user.

Tests

Run the tests using the following command:

```sh {"id":"01HHMEVXQMWZ24KMH4YBAF8H1M"}
go test

```

Cleanup

To stop and remove the containers, run:

```sh {"id":"01HHMEVXQMWZ24KMH4YBB196SP"}
docker-compose down

```
