package middlewares

import (
	"net/http"
	"strings"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
)

func ParseJWTMiddlware(c *gin.Context) {
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
	_, err := utils.VerifyJWTCustom(slice[1], &jwt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
		return
	}
	if datastore.DB.GetUserByID(jwt.UserID) == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
		return
	}

	c.Set("user_id", jwt.UserID)
	c.Next()
}

func AuthMiddleware(c *gin.Context) {
	if c.GetString("user_id") == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
