package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.String(http.StatusInternalServerError, "Oops, an unexpected error occurred, please try again")
				c.Abort()
			}
		}()
		c.Next()
	}
}
