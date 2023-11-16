package addr

import (
	"gbv2/config/log"
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
)

var AddrDB *geoip2.DBReader

func InitAddrDB() {
	statik, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatalw("ip地址数据加载失败", err)
	}
	AddrDB = statik
}
