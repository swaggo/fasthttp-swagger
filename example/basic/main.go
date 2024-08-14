package main

import (
	fastHttpSwagger "github.com/swaggo/fasthttp-swagger"
	"github.com/swaggo/fasthttp-swagger/example/basic/api"
	"github.com/valyala/fasthttp"
	"strings"

	_ "github.com/swaggo/fasthttp-swagger/example/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {
	fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.RequestURI())
		switch {
		case strings.HasPrefix(path, "/swagger"):
			fastHttpSwagger.WrapHandler(fastHttpSwagger.InstanceName("swagger"))(ctx)
		case strings.HasPrefix(path, "/testapi/get-string-by-int/"):
			api.GetStringByInt(ctx)
		case strings.HasPrefix(path, "/testapi/get-struct-array-by-string/"):
			api.GetStructArrayByString(ctx)
		}
	})
}
