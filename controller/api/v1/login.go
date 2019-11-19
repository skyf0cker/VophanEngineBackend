package v1

import (
	"VophanEngineBackend/model"
	"VophanEngineBackend/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	user, err := model.GetUserByName(username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "200",
			"msg":    "error",
			"error":  err.Error(),
		})
	} else {
		if password == user.Password {
			token, jwt_err := jwt.GenerateToken(username, password)
			if jwt_err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status": "500",
					"msg":    "server Error",
				})
			}
			ctx.SetCookie("jwt_token", token, 3600*24*3, "/", "localhost", false, true)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "200",
				"msg":    "success",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "200",
				"msg":    "密码错误",
			})
		}
	}
}
