package middlewares

import (
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidateHeaders 用于API接口POST, PUT检查请求头
func ValidateHeaders(ctx *gin.Context) {
	method := ctx.Request.Method
	if method == "POST" || method == "PUT" {
		contentType := strings.TrimSpace(ctx.GetHeader("Content-Type"))
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil || mediaType != "application/json" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "content-type header is not application/json"})
			return
		}
	}
	ctx.Next()
}
