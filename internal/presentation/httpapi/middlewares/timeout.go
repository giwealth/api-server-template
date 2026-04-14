package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var timeout = 5 * time.Second

func Timeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	done := make(chan struct{})
	go func() {
		c.Next()
		close(done)
	}()

	select {
	case <-ctx.Done():
		c.Abort()
		if !c.Writer.Written() {
			c.JSON(http.StatusGatewayTimeout, gin.H{"code": 50004, "message": "请求超时"})
		}
	case <-done:
	}
}
