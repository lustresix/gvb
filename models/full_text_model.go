package models

import (
	"context"
	"gbv2/config/es"
	"gbv2/config/log"
)

type FullTextModel struct {
	ID    string `json:"id"`    // es的id
	Key   string `json:"key"`   // 文章关联的id
	Title string `json:"title"` // 文章标题
	Slug  int    `json:"slug"`  // 跳转地址
	Body  string `json:"body"`  // 文章内容

}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "key":{
		"type":"keyword"
      },
      "title": { 
        "type": "text"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
	}
  }
}
`
}

// IndexExists 索引是否存在
func (a FullTextModel) IndexExists() bool {
	exists, err := es.ES.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a FullTextModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		err := a.RemoveIndex()
		if err != nil {
			return err
		}
	}
	// 没有索引
	// 创建索引
	createIndex, err := es.ES.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		log.Errorw("创建索引失败", "err", err)
		return err
	}
	if !createIndex.Acknowledged {
		log.Errorw("创建失败", "err", err)
		return err
	}
	log.Infow("索引创建成功", "索引为", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a FullTextModel) RemoveIndex() error {
	log.Infow("索引存在，删除索引")
	// 删除索引
	indexDelete, err := es.ES.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		log.Errorw("删除索引失败", "err", err)
		return err
	}
	if !indexDelete.Acknowledged {
		log.Errorw("删除索引失败", "err", err)
		return err
	}
	context.Background()
	log.Infow("索引删除成功")
	return nil
}

// Create 添加的方法
