package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateHeaders 用于API接口POST, PUT检查请求头
func ValidateHeaders(ctx *gin.Context) {
	headers := ctx.Request.Header
	method := ctx.Request.Method
	if method == "POST" || method == "PUT" {
		t := headers["Content-Type"]
		if len(t) == 0 || t[0] != "application/json" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "content-type header is not application/json"})
			return
		}
	}
	ctx.Next()
}
