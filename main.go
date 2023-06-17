package main

import (
	"ecs-mock/handler"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func main() {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		slog.Info(ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
	})
	r.GET("/login", handler.Login)
	ns := r.Group("/object")
	ns.Use(getAuthMiddleware())
	ns.GET("/namespaces.json", handler.ListNs)
	ns.GET("/users/:item", handler.ListNativeUsers)
	ns.GET("/users/:item/info.json", handler.GetNativeUser)
	r.GET("/iam", handler.IAMAction)
	r.POST("/iam", handler.IAMAction)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader(handler.AuthHeader)
		if auth != handler.MockToken {
			ctx.AbortWithStatus(401)
			return
		}
		ctx.Next()
	}
}
