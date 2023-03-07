package ctype

import "github.com/goccy/go-json"

type ImageType int

const (
	Local ImageType = 1 // QQ
	Qiniu ImageType = 2 // github
)

func (i ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i ImageType) String() string {
	var str string
	switch i {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	}
	return str
}
