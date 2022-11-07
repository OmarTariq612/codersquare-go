package middlewares

import (
	"net/http"
	"strings"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
)

func ParseJWTMiddleware(db datastore.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}
		slice := strings.Split(authHeader, " ")
		if len(slice) != 2 {
			c.Next()
			return
		}
		jwt := types.JWTObject{}
		_, err, isExpired := utils.VerifyJWTCustom(slice[1], &jwt)
		if err != nil {
			if isExpired {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": types.ErrTokenExpired})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": types.ErrBadToken})
			return
		}
		if db.GetUserByID(jwt.UserID) == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": types.ErrUserNotFound})
			return
		}

		c.Set("user_id", jwt.UserID)
		c.Next()
	}
}

func AuthMiddleware(c *gin.Context) {
	if c.GetString("user_id") == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
