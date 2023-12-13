##CRUD in Go with GORM, Docker Compose, Redis, and PostgreSQL

This is a simple example of a CRUD application in Go that uses GORM as the ORM, Docker Compose for container orchestration, and Redis and PostgreSQL as databases.
Configuration

Make sure you have Docker and Docker Compose installed on your machine.

    Clone the repository:
    
    bash

```sh
git clone github.com/isaquerr25/crud-go-gin-redis-gorm-testify-docker
```

Create a .env file in the project root and configure the necessary environment variables. An example can be found in the .env.example file.

Start the containers using Docker Compose:

bash

```sh
docker-compose up -d
```

Run the application:

bash

```sh
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

bash

```sh
go test
```

Cleanup

To stop and remove the containers, run:

bash

```sh
docker-compose down
```

Contributing

Feel free to contribute to this project! Just follow the contribution guidelines and submit your pull requests.
