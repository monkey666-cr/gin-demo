package middleware

import (
	"github.com/chenrun666/gin_demo/handler"
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/chenrun666/gin_demo/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
