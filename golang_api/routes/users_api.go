package routes

import (
	"net/http"
	"udemy-multi-api-golang/models"

	"github.com/gin-gonic/gin"
)

func sign_up(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse the value.", "user": user})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Saving User in the Database", "Error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "user Created",
		"event": user})
}

func user_Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not Login.", "error": err.Error()})
	}

}
