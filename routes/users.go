package routes

import (
	"net/http"

	"example.com/go-rest/models"
	"example.com/go-rest/utils"
	"github.com/gin-gonic/gin"
)

const (
	msgSaveUserFailed = "Could not save user."
	msgUserCreated    = "User created!"
	msgAuthFailed     = "Could not authenticate user."
	msgUserLoggedIn   = "User Logged in!"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseRequestData})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgSaveUserFailed})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": msgUserCreated})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseRequestData})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgAuthFailed})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": msgUserLoggedIn, "token": token})

}
