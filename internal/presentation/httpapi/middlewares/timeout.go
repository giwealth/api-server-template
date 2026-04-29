package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

var timeout = 5 * time.Second

// Timeout 超时中间件
func Timeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
