package test

import (
	"VophanEngineBackend/utils/elastic"
	"context"
	"encoding/json"
	"fmt"
	ela "gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestElastic_create(t *testing.T) {
	article := elastic.Articles{
		Title:"testtitle",
		Time:"testtime",
		Content:"testcontent",
	}

	put, err := elastic.Create(article)
	fmt.Println(put.Result)
	if err != nil{
		fmt.Println(err.Error())
	}
}

func TestElastic_update(t *testing.T) {
	q := ela.NewQueryStringQuery("title:testtitle")
	res, err := elastic.Client.Search("articles").Type("_doc").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)
		for _, hit := range res.Hits.Hits {
			var t elastic.Articles
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}
			fmt.Printf("article name %s : %s\n", t.Title, t.Time)
			//ures, uerr := elastic.Update(hit.Id, map[string]interface{}{"time":"9102"})
			//if uerr == nil{
			//	fmt.Println(ures.Result)
			//}
		}
	}
}

func TestElastic_Delete(t *testing.T) {
	q := ela.NewQueryStringQuery("title:testtitle")
	res, err := elastic.Client.Search("articles").Type("_doc").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)
		for _, hit := range res.Hits.Hits {
			var t elastic.Articles
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}
			fmt.Printf("article name %s : %s\n", t.Title, t.Time)
			//dres, derr := elastic.Delete(hit.Id)
			//if derr == nil{
			//	fmt.Println(dres.Result)
			//}
		}
	}
}

func TestElastic_Query(t *testing.T) {
	res, err := elastic.Query("人工智能")

	if err == nil{
		var articles []elastic.Articles
		for _, hit := range res.Hits.Hits {
			var t elastic.Articles
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			articles = append(articles, t)
			if err != nil {
				fmt.Println("Deserialization failed")
			}
			fmt.Printf("article name %s : %s\n", t.Title, t.Content)
		}
		fmt.Println(articles)
	}
}