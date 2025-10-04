package middlewares

import (
	"net/http"
	"strings"

	"example.com/go-rest/constants"
	"example.com/go-rest/utils"
	"github.com/gin-gonic/gin"
)

const msgNotAuthorized = "Not authorized."

func Authenticate(context *gin.Context) {
	auth := context.Request.Header.Get("Authorization")

	if auth == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": msgNotAuthorized})
		return
	}

	// Sanitize the received token
	token := strings.TrimSpace(auth)
	token = strings.Trim(token, `"'`)

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": msgNotAuthorized})
		return
	}

	context.Set(constants.ContextUserID, userId)

	context.Next()
}
