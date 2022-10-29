package middlewares

import (
	"net/http"
	"strings"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		slice := strings.Split(authHeader, " ")
		if len(slice) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		jwtObj := types.JWTObject{}
		_, err := utils.VerifyJWTCustom(slice[1], &jwtObj)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
			return
		}
		if datastore.DB.GetUserByID(jwtObj.UserID) == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
			return
		}

		ctx.Set("user_id", jwtObj.UserID)
		ctx.Next()
	}
}
