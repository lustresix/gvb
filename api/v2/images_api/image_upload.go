package images_api

import (
	"gbv2/config/log"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/fs"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImagesApi) ImageUploadView(c *gin.Context) {
	var imagePath = viper.GetString("upload.path")
	var maxSize = viper.GetFloat64("upload.size")

	// 判断保存路径是否存在
	_, err := os.ReadDir(imagePath)
	if err != nil {
		_ = os.MkdirAll(imagePath, fs.ModePerm)
	}
	// 获取文件
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("文件不存在", c)
		return
	}

	var resList []FileUploadResponse
	// 循环单独上传
	for _, file := range fileList {
		// TODO: 限制文件类型

		size := float64(file.Size) / float64(1024*1024)
		// 判断大小是否超过
		if size > maxSize {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "啊啊太大了！",
			})
			continue

		}
		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功！",
		})
		// 更名保存
		dst := rename(file, imagePath)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "上传错误！",
			})
			continue
		}

		log.Infow("上传图片成功", "place: ", dst)
	}
	res.OKWithData(resList, c)
}

// 防止在图片名称上动手脚
func rename(file *multipart.FileHeader, imagePath string) string {
	//获取图像后缀
	y := path.Ext(file.Filename)
	//获取时间戳防止重复 !需要精准到纳秒，防止传输过快产生同名，然后出错
	a := time.Now().UnixNano()
	//获取一个1w以内的随机数
	b := rand.Intn(10000)
	//将时间辍的类型转换
	z := strconv.FormatInt(a, 10)
	//将随机数转换类型
	x := strconv.FormatInt(int64(b), 10)
	//写入保存位置与自定义名称，并且带上文件自带后缀名
	dst := path.Join(imagePath, z+x+y)
	return dst
}
