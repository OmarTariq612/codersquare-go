package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.String(http.StatusInternalServerError, "Oops, an unexpected error occurred, please try again")
			}
		}()
		ctx.Next()
	}
}
