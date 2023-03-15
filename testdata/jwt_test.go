package testdata

import (
	"fmt"
	"gbv2/utils/jwt"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := jwt.GetToken(jwt.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "linxx",
		NickName: "xxx",
	})
	fmt.Println(token, err)
	parseToken, err := jwt.ParseToken(token)
	fmt.Println(parseToken, err)
}
