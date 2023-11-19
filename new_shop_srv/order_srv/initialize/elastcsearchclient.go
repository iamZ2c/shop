package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

type AccountAddr struct {
	FirstName string `json:"first_name"`
	Address   string `json:"address"`
}

func InitElasticSearchClient() (err error) {
	var (
		client *elastic.Client
		//indexResp *elastic.IndexResponse
		searchRes *elastic.SearchResult
		logger    *log.Logger
		query     *elastic.MatchQuery
		data      []byte
	)
	logger = log.New(os.Stdout, "test", log.LstdFlags)
	client, err = elastic.NewClient(
		elastic.SetURL("http://192.168.10.6:9200"),
		elastic.SetSniff(false),
		elastic.SetTraceLog(logger),
	)
	a := AccountAddr{}

	query = elastic.NewMatchQuery("address", "cq")
	if searchRes, err = client.Search().Query(query).Do(context.Background()); err != nil {
		panic(err)
	}
	//total := searchRes.Hits.TotalHits.Value
	for _, v := range searchRes.Hits.Hits {
		if data, err = v.Source.MarshalJSON(); err != nil {
			panic(err)
		}
		json.Unmarshal(data, &a)
		fmt.Println(a)

	}

	// 插入数据
	//if indexResp, err = client.Index().Index("test").BodyJson(a).Do(context.Background()); err != nil {
	//	panic(err)
	//}
	//fmt.Println(indexResp.Id, indexResp.Index, indexResp.Status)
	return
}
