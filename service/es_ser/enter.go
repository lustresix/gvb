package es_ser

import (
	"context"
	"errors"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"github.com/goccy/go-json"
	"github.com/olivere/elastic/v7"
)

func CommList(key string, page, limit int) (list []models.ArticleModel, count int, err error) {

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
	do, err := es.ES.Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).Size(limit).
		Do(context.Background())
	if err != nil {
		log.Errorw("err", "err", err)
	}
	count = int(do.Hits.TotalHits.Value)
	demolist := []models.ArticleModel{}
	for _, hit := range do.Hits.Hits {
		var model models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			log.Errorw("err", "err", err)
			continue
		}
		err = json.Unmarshal(data, &model)
		if err != nil {
			log.Errorw("err", "err", err)
			continue
		}
		model.ID = hit.Id
		demolist = append(demolist, model)
	}
	return demolist, count, err
}

func CommDetail(id string) (model models.ArticleModel, err error) {
	res, err := es.ES.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	return
}

func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := es.ES.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}
