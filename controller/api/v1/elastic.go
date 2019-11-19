package v1

import (
	"VophanEngineBackend/utils/elastic"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func UploadData(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":"500",
			"msg": "服务器内部错误",
		})
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		article := elastic.Articles{}
		jerr := json.Unmarshal(message, &article)
		if jerr != nil{
			log.Println(jerr)
			continue
		}
		res, elastic_err := elastic.Create(article)
		if elastic_err != nil{
			break
		} else {
			log.Println(res.Result)
		}
		////写入ws数据
		//err = ws.WriteMessage(mt, message)
		//if err != nil {
		//	break
		//}
	}

}
