package models

import (
	"context"
	"encoding/json"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models/ctype"
	"github.com/olivere/elastic/v7"
)

type ArticleModel struct {
	ID        string `json:"id"`         // es的id
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间

	Title    string `json:"title"`              // 文章标题
	Keyword  string `json:"keyword,omit(list)"` // 关键字
	Abstract string `json:"abstract"`           // 文章简介
	Content  string `json:"content,omit(list)"` // 文章内容

	LookCount     int `json:"look_count"`     // 浏览量
	CommentCount  int `json:"comment_count"`  // 评论量
	DiggCount     int `json:"digg_count"`     // 点赞量
	CollectsCount int `json:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id"`        // 用户id
	UserNickName string `json:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar"`    // 用户头像

	Category string `json:"category"`          // 文章分类
	Source   string `json:"source,omit(list)"` // 文章来源
	Link     string `json:"link,omit(list)"`   // 原文链接

	BannerID  uint   `json:"banner_id"`  // 文章封面id
	BannerUrl string `json:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags"` // 文章标签
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" url:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      },
      "tags": { 
        "type": "keyword"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

func ESCreateIndex() {
	//_ = ArticleModel{}.CreateIndex()
	_ = FullTextModel{}.CreateIndex()
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
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
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
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
	log.Infow("索引创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
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
	log.Infow("索引删除成功")
	return nil
}

// Create 添加的方法
func (a *ArticleModel) Create() (err error) {
	indexResponse, err := es.ES.Index().
		Index(a.Index()).
		BodyJson(a).Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	res, err := es.ES.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		log.Errorw(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (a *ArticleModel) GetDataByID(id string) error {
	res, err := es.ES.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	err = json.Unmarshal(res.Source, a)
	return err

}
