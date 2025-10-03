package middlewares

import (
	"net/http"
	"strings"

	"example.com/go-rest/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	auth := context.Request.Header.Get("Authorization")

	if auth == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Sanitize the received token
	token := strings.TrimSpace(auth)
	token = strings.Trim(token, `"'`)

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
