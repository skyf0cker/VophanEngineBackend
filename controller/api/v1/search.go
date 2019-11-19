package v1

import (
	"VophanEngineBackend/utils/elastic"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(ctx *gin.Context) {
	if keywords := ctx.Query("keyword"); keywords != "" {
		res, err := elastic.Query(keywords)
		if err == nil {
			var articles []elastic.Articles
			for _, hit := range res.Hits.Hits {
				var t elastic.Articles
				err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
				articles = append(articles, t)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"status": "500",
						"msg":    "服务器错误",
					})
				}
			}
			ctx.HTML(http.StatusOK, "result.html", &articles)
			//ctx.JSON(http.StatusOK, gin.H{
			//	"status": "200",
			//	"msg":    "success",
			//	"data":   articles,
			//})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "500",
				"msg":    "服务器错误",
			})
		}
	}
}
