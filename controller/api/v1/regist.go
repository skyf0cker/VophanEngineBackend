package v1

import (
	"VophanEngineBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Regist(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	if err := model.CreateUser(&model.User{
		Username: username,
		Password: password,
		Email:    email,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "200",
			"msg:":   "error",
			"error":  err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "200",
			"msg:":   "success",
		})
	}

}
