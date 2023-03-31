package es_ser

import (
	"context"
	"gbv2/config/es"
	"gbv2/models"
	"github.com/olivere/elastic/v7"
)

func ArticleUpdate(id string, data map[string]any) error {
	_, err := es.ES.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	return err
}
