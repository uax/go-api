package middleware

import "github.com/gin-gonic/gin"

func CORS(ctx *gin.Context) {
	method := ctx.Request.Method

	ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	// 默认过滤这两个请求,使用204(No Content)这个特殊的http status code
	if method == "OPTIONS" || method == "HEAD" {
		ctx.AbortWithStatus(204)
		return
	}
	ctx.Next()
}
