package pwd

import (
	"gbv2/config/log"
	"golang.org/x/crypto/bcrypt"
)

func HashPwd(password string) (hashPwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Errorw("加密错误", "err", err)
		return "", err
	}
	return string(hash), nil
}

func CheckPwd(hashPwd string, pwd string) bool {
	bytes := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(bytes, []byte(pwd))
	if err != nil {
		log.Errorw("解密失败", "err", err)
		return false
	}
	return true
}
