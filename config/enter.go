package config

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/config/redis"
	"gbv2/config/system"
	"gbv2/routers"
	"github.com/spf13/viper"
)

func InitConfig() {
	log.LogInit(logOptions())
	defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中
	mysql.DBInit()
	redis.RedisInit()
	router := routers.InitRouter()
	router.Run(system.Addr())
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
