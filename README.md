# fasthttp-swagger

[fasthttp](https://github.com/valyala/fasthttp) middleware to automatically generate RESTFUL API documentation with Swagger 2.0.

[![Build Status](https://github.com/swaggo/fasthttp-swagger/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/features/actions)
[![Codecov branch](https://img.shields.io/codecov/c/github/swaggo/fasthttp-swagger/master.svg)](https://codecov.io/gh/swaggo/fasthttp-swagger)
[![Go Report Card](https://goreportcard.com/badge/github.com/swaggo/fasthttp-swagger)](https://goreportcard.com/report/github.com/swaggo/fasthttp-swagger)
[![GoDoc](https://godoc.org/github.com/swaggo/fasthttp-swagger?status.svg)](https://godoc.org/github.com/swaggo/fasthttp-swagger)
[![Release](https://img.shields.io/github/release/swaggo/fasthttp-swagger.svg?style=flat-square)](https://github.com/swaggo/fasthttp-swagger/releases)

## Usage

### Start using it

1. Add comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
go get -u github.com/swaggo/swag/cmd/swag
```

3. Run the [Swag](https://github.com/swaggo/swag) at your Go project root path(for instance `~/root/go-peoject-name`),
   [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`)
   at `~/root/go-peoject-name/docs`.

```sh
swag init
```

4. Download [fasthttp-swagger](https://github.com/swaggo/fasthttp-swagger) by using:

```sh
go get -u github.com/swaggo/fasthttp-swagger
go get -u github.com/swaggo/files
```

Import following in your code:

```go
import "github.com/swaggo/fasthttp-swagger" // fasthttp-swagger middleware

```

### Canonical example:

Now assume you have implemented a simple api as following:

```go
// A get function which returns a hello world string by json
func Helloworld(ctx *fasthttp.RequestCtx)  {
    ctx.SetStatusCode(http.StatusOK)
    ctx.Write([]byte("helloworld"))
}

```

So how to use fasthttp-swagger on api above? Just follow the following guide.

1. Add Comments for apis and main function with fasthttp-swagger rules like following:

```go
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(ctx *fasthttp.RequestCtx)  {
    ctx.SetStatusCode(http.StatusOK)
    ctx.Write([]byte("helloworld"))
}
```

2. Use `swag init` command to generate a docs, docs generated will be stored at
3. import the docs like this:
   I assume your project named `github.com/go-project-name/docs`.

```go
import (
   docs "github.com/go-project-name/docs"
)
```

4. build your application and after that, go to http://localhost:8080/swagger/index.html ,you to see your Swagger UI.

5. The full code and folder relatives here:

```go
package main

import (
   _ "github.com/go-project-name/docs"
   fasthttpSwagger "github.com/swaggo/fasthttp-swagger"
)
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(ctx *fasthttp.RequestCtx)  {
   ctx.SetStatusCode(http.StatusOK)
   ctx.Write([]byte("helloworld"))
}

func main()  {
   fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
      path := string(ctx.RequestURI())
      switch {
      case strings.HasPrefix(path, "/swagger"):
         fastHttpSwagger.WrapHandler(fastHttpSwagger.InstanceName("swagger"))(ctx)
      case path == "/api/v1/example/helloworld":
         Helloworld(ctx)
      }
   })
}
```

Demo project tree, `swag init` is run at relative `.`

```
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go
```