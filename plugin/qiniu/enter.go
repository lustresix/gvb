package qiniu

import (
	"bytes"
	"context"
	"errors"
	"gbv2/config/qiniu"
	"gbv2/utils"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

func qiNiuOptions() *qiniu.QiNiu {
	return &qiniu.QiNiu{
		Enable:    viper.GetBool("qi_niu.enable"),
		AccessKey: viper.GetString("qi_niu.access_key"),
		SecretKey: viper.GetString("qi_niu.secret_key"),
		Bucket:    viper.GetString("qi_niu.bucket"),
		Zone:      viper.GetString("qi_niu.zone"),
		CDN:       viper.GetString("qi_niu.cdn"),
		Size:      viper.GetFloat64("qi_niu.size"),
	}
}

func getToken() string {
	q := qiNiuOptions()
	PutPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := PutPolicy.UploadToken(mac)
	return upToken
}

func getCfg() storage.Config {
	q := qiNiuOptions()
	cfg := storage.Config{}
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	return cfg
}

func UploadImage(data []byte, filename string) (filepath string, rename string, err error) {
	q := qiNiuOptions()
	if !q.Enable {
		return "", "", errors.New("七牛还没打开啊")
	}
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", "", errors.New("请配置access_key和secret_key")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", "", errors.New("啊啊太大了不要塞进来啊")
	}
	upToken := getToken()
	cfg := getCfg()

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	_, key := utils.Rename(filename, "")

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", "", err
	}
	return q.CDN, ret.Key, nil
}
