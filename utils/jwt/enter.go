package jwt

import (
	"errors"
	"gbv2/config/log"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
	"time"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	Username string `json:"username"`  // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var (
	Secret   = viper.GetString("jwt.secret")
	Expires  = viper.GetDuration("jwt.expires")
	Issuer   = viper.GetString("jwt.issuer")
	MySecret = []byte(Secret)
)

// GetToken 创建 Token
func GetToken(user JwtPayLoad) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * Expires)), // 默认2小时过期
			Issuer:    Issuer,                                      // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Errorw("token parse err:", "err", err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
