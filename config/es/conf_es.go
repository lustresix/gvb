package es

import (
	"gbv2/config/log"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var ES *elastic.Client

func EsConnect() {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := viper.GetString("es.host")
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		log.Errorw("es连接失败 %s", err.Error())
	}
	log.Infow("Elasticsearch连接成功")
	ES = c
}
