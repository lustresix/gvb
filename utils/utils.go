package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"path"
	"strconv"
	"time"
)

// InList 判断是否在列表里
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Rename 防止在图片名称上动手脚
func Rename(filename string, imagePath string) (dst string, name string) {
	//获取图像后缀
	y := path.Ext(filename)
	//获取时间戳防止重复 !需要精准到纳秒，防止传输过快产生同名，然后出错
	a := time.Now().UnixNano()
	//获取一个1w以内的随机数
	b := rand.Intn(10000)
	//将时间辍的类型转换
	z := strconv.FormatInt(a, 10)
	//将随机数转换类型
	x := strconv.FormatInt(int64(b), 10)
	//写入保存位置与自定义名称，并且带上文件自带后缀名
	dst = path.Join(imagePath, z+x+y)
	name = z + x + y
	return dst, name
}

// MD5 MD5加密
func MD5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
