package utils

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
)

var db *geoip2.DBReader

func init() {
	db, _ = geoip2db.NewGeoipDbByStatik()
	defer db.Close()
}
