package system

import (
	"fmt"
	"gbv2/config/log"
	"github.com/spf13/viper"
)

func Addr() string {
	s := fmt.Sprintf("%v:%v", viper.GetString("system.host"), viper.GetString("system.port"))
	log.Infow("blog sever is running at " + s)
	return s
}
