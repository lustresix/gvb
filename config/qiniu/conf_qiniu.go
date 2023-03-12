package qiniu

type QiNiu struct {
	Enable    bool    `json:"enable"`
	CDN       string  `json:"cdn"`
	AccessKey string  `json:"accessKey"`
	SecretKey string  `json:"secretKey"`
	Bucket    string  `json:"bucket"`
	Zone      string  `json:"zone"`
	Size      float64 `json:"size"`
}
