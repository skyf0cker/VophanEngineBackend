package v1

import (
	v1 "VophanEngineBackend/controller/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.LoadHTMLGlob("web/*")
	r.Static("static", "./static")
	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	//r.Use(middleware.Logger())

		r.GET("index", v1.Index)
	r.POST("api/v1/regist", v1.Regist)
	r.GET("api/v1/login", v1.Login)
	r.GET("api/v1/elastic", v1.UploadData)

	api := r.Group("api/v1")
	//api.Use(middleware.JWT())
	{
		api.GET("/search", v1.Search)
	}

	return r
}

