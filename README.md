# The Bank

A example of bank transaction, made with GOlang, Gin Framework, Restful and GraphQL.

- [Endpoints](#Endpoints)
- [Insomnia](#Insomnia)
- [Get Started](#Get-Started)
- [Get Started - without docker](##Get-Stared-without-docker)
- [Run Debug with VScode and DB in docker](##Run-Debugg)
- [Unity Test and Test Integration](##Unity-Test-and-Test-Integration)
- [Step By Step](#Step-By-Step)
- [Commands](#Commands)
- [Technologies](#Technologies)
- [References](#References)

## Endpoints

| Endpoint              | Method | Action  |
| --------------------- | ------ | ------- |
| /graphql              | POST   |         |
| /graphiql             | GET    |         |
| /accounts             | GET    | Index   |
| /accounts             | POST   | Store   |
| /accounts             | PATCH  | Update  |
| /accounts/:id         | DELETE | Destroy |
| /accounts/:id         | GET    | Show    |
| /transactions         | GET    | Index   |
| /transactions         | POST   | Store   |
| /transactions         | PATCH  | Update  |
| /transactions/:id     | DELETE | Destroy |
| /transactions/:id     | GET    | Show    |
| /operations-types     | GET    | Index   |
| /operations-types     | POST   | Store   |
| /operations-types     | PATCH  | Update  |
| /operations-types/:id | DELETE | Destroy |
| /operations-types/:id | GET    | Show    |

## Insomnia

If you use Insomnia, just import a [insomnia.yaml](insomnia.yaml)

## Get Started

Just one command

```bash
docker-compose up --build
```

> Open on you best browser: http://localhost:8085 <br />
> GraphQL Playground: http://localhost:8085/gaphiql<br />
> To see a database open adminer on: http://localhost:8086 <br />
> Configurations for connect on database see: [/docker/local.env](/docker/local.env)
> IMPORTANT! If you not have a docker see: [Get Started - without docker](#Get-Stared-without-docker)

## Structure

```
├── config - (configuration)
├── controllers (RESTfull methods)
├── core (Base methos for controllers, models and services)
├── database - (database configuration, migrate and seed data)
├── graphql
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
├── models (ORM models based on Database tables)
├── routes (files with routes)
├── services (with business rules and intermediate a database)controlling the generated code.
├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for
└── main.go                  - The entry point to your app. Customize it however you see fit
```

## Commands

Enter on docker container to exec any command.

```bash
docker exec -it labbankgo-api /bin/bash
```

To generate graphql files.

```bash
gqlgen generate
```

> On docker

```bash
~/go/bin/gqlgen generate
```

> Out of docker

#### Get Stared without docker

> A Database PostgreSQL is needed, configurations for it in docker/local.env

```bash
gin --port=8080 #or ~/go/bin/gin --port=8080
```

## Run Debugg

> With VScode and SGBD on Docker

```bash
docker-compose -f docker-compose-db.yml up
```

## Unity Test and Test Integration

```bash
go test -v ./src/tests/
```

With code coverage

```bash
go test -cover -coverprofile=c.out ./src/tests/
go tool cover -html=c.out -o coverage.html
```

## Step By Step (for create this project)

Create a mod file for dependence management

```bash
go mod init github.com/ruyjfs/lab-bank-go
```

Install Gin Framework

```bash
go get github.com/codegangsta/gin
```

Create a main.go with this example code

```bash
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		c.String(http.StatusOK, "Hello World", name)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

Install GQLgen for GraphQL

```bash
go get github.com/99designs/gqlgen
```

Initial structure for GraphQL

```bash
go run github.com/99designs/gqlgen init
```

Install GORM

```bash
go get -u gorm.io/gorm && \
go get -u gorm.io/driver/postgres
```

Start project with Hot Reload

```bash
~/go/bin/gin
```

Default start project

```bash
go run main.go
```

> Open on you best browser: http://localhost:8085

## Technologies:

- [Golang (Language)](https://golang.org)
- [Gin (Framework)](https://gin-gonic.com/docs/)
- [GORM (ORM)](https://gorm.io/docs/index.html)
- [Gqlgen (GraphQL)](https://gqlgen.com)
- [PostgreSQL (SGBD)](https://www.postgresql.org/docs/online-resources/)
- [Docker Compose (Local Environment)](https://docs.docker.com/compose/compose-file/)

## References

- [Effective GO](https://golang.org/doc/effective_go.html)
- [JsonAPI (Restful API Specification)](https://jsonapi.org)
- [Style Guide - GOlang Code Review](https://github.com/golang/go/wiki/CodeReviewComments)
- [Style Guide - Google](https://google.github.io/styleguide/)
- [Style Guide - Uber GO](https://github.com/uber-go/guide/blob/master/style.md/)
- [Style Guide - Source Graph](https://about.sourcegraph.com/handbook/engineering/languages/go/)
- [How to GraphQL with GO](https://www.howtographql.com/graphql-go/0-introduction/)
- [Examples with GO](https://gobyexample.com/)
