package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claims struct {
	UnionID string `json:"unionid"`
	jwt.StandardClaims
}

var jwtSecret = []byte("fe68b7a5-8f3f-4362-ab79-a0e12c717943")

// Authorize 鉴权
func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var msg string
		token := ctx.Request.Header.Get("Authorization")
		if token != "" {
			cl, err := parseToken(token)
			if err != nil {
				msg = fmt.Sprint(err)
			} else if time.Now().Unix() > cl.ExpiresAt {
				msg = "token timeout"
			}
		} else {
			msg = "token is empty"
		}

		if msg != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			ctx.Abort()
			return
		}

		ctx.Set("UserID", ctx.Request.Header.Get("user_id"))
		ctx.Next()
	}
}

// parseToken 解析token
func parseToken(token string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
