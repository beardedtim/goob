package http

import (
	"fmt"
	"mckp/goob/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func GetUserById(ctx *gin.Context) {
	userModel := data.NewUserModel()

	user, err := userModel.ById(ctx.Param("id"))

	if err != nil {
		if err.Error() == "Not Found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Cannot find User %s. Please modify your request and try again.", ctx.Param("id")),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "There has been an issue internally. Please try your request again later.",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
