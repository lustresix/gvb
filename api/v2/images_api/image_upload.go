package images_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/global"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"io/fs"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImagesApi) ImageUploadView(c *gin.Context) {
	var imagePath = viper.GetString("upload.path")
	var maxSize = viper.GetFloat64("upload.size")
	var imageModel models.ImageModel

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

		// 限制文件类型
		flag := utils.InList(path.Ext(file.Filename), global.WhiteImageList)
		if !flag {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "这个不能塞进来啊！",
			})
			continue
		}

		// 判断大小是否超过
		size := float64(file.Size) / float64(1024*1024)
		if size > maxSize {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "啊啊太大了！",
			})
			continue
		}

		// 判断是否有重复
		open, _ := file.Open()
		bytes, _ := io.ReadAll(open)
		hash := utils.MD5(bytes)
		err := mysql.DB.Take(&imageModel, "hash = ?", hash).Error
		if err == nil {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "已经在里面了！",
			})
			continue
		}

		// 更名保存
		dst, name := utils.Rename(file, imagePath)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "上传错误！",
			})
			continue
		}

		// 持久化数据库
		mysql.DB.Create(&models.ImageModel{
			Path: imagePath,
			Hash: hash,
			Name: name,
		})

		// 标记为成功
		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功！",
		})

		log.Infow("上传图片成功", "place: ", dst)
	}
	res.OKWithData(resList, c)
}
