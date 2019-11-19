package elastic

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"strconv"
)

var Client *elastic.Client
var host = "http://127.0.0.1:9200"

type Articles struct {
	Title   string `json:"title"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error

	Client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := Client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := Client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

func Create(art Articles) (*elastic.IndexResponse, error) {
	put, err := Client.Index().
		Index("articles").
		Type("_doc").
		BodyJson(art).
		Do(context.Background())
	return put, err
}

func Delete(id string) (*elastic.DeleteResponse, error) {
	res, err := Client.Delete().Index("articles").
		Type("_doc").
		Id(id).
		Do(context.Background())
	return res, err
}

func Update(id string, change map[string]interface{}) (*elastic.UpdateResponse, error) {
	res, err := Client.Update().
		Index("articles").
		Type("_doc").
		Id(id).
		Doc(change).
		Do(context.Background())
	return res, err
}

func Gets(id int) (*elastic.GetResult, error) {
	//通过id查找
	get1, err := Client.Get().Index("articles").Type("_doc").Id(strconv.Itoa(id)).Do(context.Background())
	return get1, err
}

func Query(val string) (*elastic.SearchResult, error) {
	var res *elastic.SearchResult
	var err error

	matchPhraseQuery := elastic.NewMatchPhraseQuery("content", val)
	res, err = Client.Search("articles").Type("_doc").Query(matchPhraseQuery).Do(context.Background())
	return res, err
}
