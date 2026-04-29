package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var timeout = 5 * time.Second

// Timeout 超时中间件
func Timeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	done := make(chan struct{}, 1)
	go func() {
		c.Next()
		done <- struct{}{}
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		SaveRequestError(c, context.DeadlineExceeded)
		c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
			"code":    50004,
			"message": "请求超时",
		})
		return
	}
}
