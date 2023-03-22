package es

import (
	"context"
	"encoding/json"
	"fmt"
	"gbv2/config/log"
	"github.com/olivere/elastic/v7"
)

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

// Create 创建
func Create(data *DemoModel) (err error) {
	indexResponse, err := ES.Index().
		Index(data.Index()).
		BodyJson(data).Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return err
	}
	data.ID = indexResponse.Id
	return nil
}

// FindList 列表查询
func FindList(key string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := ES.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			log.Errorw(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			log.Errorw("err", "err", err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

func FindSourceList(key string, page, limit int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := ES.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Source(`{"_source": ["title"]}`).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []DemoModel{}
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			log.Errorw(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			log.Errorw("err", "err", err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	fmt.Println(demoList, count)
}

// Update 更新
func Update(id string, data *DemoModel) error {
	_, err := ES.
		Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return err
	}
	log.Infow("更新demo成功")
	return nil
}

// Remove 批量删除
func Remove(idList []string) (count int, err error) {
	bulkService := ES.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}
