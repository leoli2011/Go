package main

import (
	"fmt"
	"log"
	"os"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"reflect"
	"encoding/json"
)

var host = "http://138.68.48.23:9200"
var client *elastic.Client

type Employee struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	About string `json:"about"`
	Interests []string `json:"interests"`
}
func init() {
	errorlog := log.New(os.Stdout, "es", log.LstdFlags)
	var err error
	client,err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("es return with code %d and version %s\n", code, info.Version.Number)
	esversionCode, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ES version %s\n", esversionCode)
}

func create() {
	e1 := Employee{"leo", "Li", 18, "i like movie", []string{"movie"}}
	//e1 := `{"firstname":"john", "lastname":"smith", "age":22, "about":"i like play computer", "interests":["computer","music"]}`
	put, err := client.Index().Index("info").Type("employee").Id("3").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("indexed %s, to index %s, type %s\n", put.Id, put.Index, put.Type)
}

func get() {
	get, err := client.Get().Index("info").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Printf("got document %s in version %d from index %s, type %s\n",
					get.Id, get.Version, get.Index, get.Type)
	}
}

func update() {
	res, err:= client.Update().Index("info").Type("employee").Id("2").Doc(map[string]interface{}{"age":88}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("update age %s", res.Result)
}

func delete() {
	res, err := client.Delete().Index("info").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("delete result %s", res.Result)
}

func query() {
	var res *elastic.SearchResult
	var err error
	res, err = client.Search("info").Type("employee").Do(context.Background())
	printEmployee(res, err)
}

func query1() {
	var res *elastic.SearchResult
	var err error
	q := elastic.NewQueryStringQuery("firstname:john")
	res, err = client.Search("info").Type("employee").Query(q).Do(context.Background())
	printEmployee(res, err)
	fmt.Println("==========")

	if res.Hits.TotalHits > 0 {
		fmt.Printf("found a total of %d Employee\n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {
			var t Employee
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				fmt.Printf("failed")
			}else {
				fmt.Printf("employee name %s:%s\n", t.Firstname, t.Lastname )
			}
		}
	} else {
		fmt.Printf("found no employee\n")
	}
}

func query3() {
	var res *elastic.SearchResult
	var err error
	boolq := elastic.NewBoolQuery()
	boolq.Must(elastic.NewMatchQuery("lastname", "Li"))
	boolq.Filter(elastic.NewRangeQuery("age").Lt(20))
	res, err = client.Search("info").Type("employee").Query(boolq).Do(context.Background())
	printEmployee(res, err)
}

func query4() {
	var res *elastic.SearchResult
	var err error
	matchPhrase := elastic.NewMatchPhraseQuery("about", "movie")
	res, err = client.Search("info").Type("employee").Query(matchPhrase).Do(context.Background())
	printEmployee(res, err)
}

func aggs() {
	var res *elastic.SearchResult
	var err error
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("info").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)
}

//分页，确定每页显示几个
func list(size, page int)  {
	var res *elastic.SearchResult
	var err error
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err = client.Search("info").Type("employee").Size(size).From((page-1)*size).Do(context.Background())
	printEmployee(res, err)
}

func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Employee)
		fmt.Printf("employer info: %#v\n", t)
	}
}

func main() {
	fmt.Println("222")
	create()
/*	fmt.Println("222")
	get()
	fmt.Printf("333\n")
	update()
	fmt.Printf("444\n")
	delete()*/
	//query()
	//query1()
	//query3()
	//query4()
	//aggs()
	list(2,1)
}
