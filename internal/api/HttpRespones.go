package api

import "github.com/kataras/iris/v12"

func NotFound(ctx iris.Context) {
	ctx.View("web/views/errors/404.html")
}

func InternalServerError(ctx iris.Context) {
	ctx.View("web/views/errors/500.html")
}
