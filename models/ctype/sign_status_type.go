package ctype

import "github.com/goccy/go-json"

type SignStatus int

const (
	SignQQ     SignStatus = 1 // QQ
	SignGithub SignStatus = 2 // github
	SignEmail  SignStatus = 3 // emailSignQQ
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGithub:
		str = "github"
	case SignEmail:
		str = "email"

	}
	return str
}
