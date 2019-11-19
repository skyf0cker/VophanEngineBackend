package middleware

import (
	"VophanEngineBackend/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code string
		var data interface{}

		//	设置code情况
		//	1 ： 成功
		// `2 ： 权限未通过
		//	3 : 鉴权失败
		//	4 : 鉴权超时
		//	5 : 非法参数

		code = "成功"

		token, perr := context.Cookie("jwt_token")

		if perr != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "401",
				"msg":    "身份过期，请重新登录",
			})
			context.Abort()
			return
		}

		if token == "" {
			code = "非法参数"
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = "权限未通过"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = "鉴权超时"
			} else {
				code = "成功"
			}
		}

		if code != "成功" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    code,
				"data":   data,
			})

			context.Abort()
			return
		}
		context.Next()
	}
}
