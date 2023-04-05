package es_ser

import (
	"context"
	"fmt"
	"gbv2/config/es"
	"gbv2/config/log"
	"gbv2/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
	Key   string `json:"key"`
}

// AsyncArticleByText 同步es数据到全文搜索
func AsyncArticleByText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)

	bulk := es.ES.Bulk()
	for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		log.Errorw(err.Error(), "err", err)
		return
	}
	log.Infow(title, fmt.Sprintf("添加成功共 %d 条！", len(result.Succeeded())))
	return
}

// DeleteFullTextByArticleID 删除文章数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := es.ES.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	log.Infow("成功删除", "删除", fmt.Sprintf(" %d 条记录", res.Deleted))
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			//if strings.TrimSpace(body) != "" {
			bodyList = append(bodyList, getBody(body))
			//}
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	return searchDataList
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
