# A Bank transaction made with GOlang and Gin Framework

### Get Started (just one command)

```bash
docker-compose up --build
```

> Open on you best browser: http://localhst:3000

Without docker<br />

> A Database PostgreSQL is needed!

```bash
~/go/bin/gin
```

> Open on you best browser: http://localhst:3000

## Commands

To generate graphql files.

```bash
~/go/bin/gqlgen generate
```

### Technologies:

- Golang (Language)
- Gin (Framework)
- Gqlgen (GraphQL)
- PostgreSQL (SGBD)
-

## Structure

|--database - (database configuration)<br />
|--graphql - (files for graphql)<br />
|--|--generated (this file is generated automatic, DO NOT EDIT)<br />
|--|--model (this file is generated automatic, DO NOT EDIT)<br />
|--routes (files with routes)<br />
|--controllers (RESTfull methods)<br />
|--services (with business rules and intermediate a database)<br />

<!-- ├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit -->

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

Start project with Hot Reload

```bash
~/go/bin/gin
```

Default start project

```bash
go run main.go
```
