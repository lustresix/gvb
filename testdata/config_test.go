package testdata

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println(viper.GetString("site_info"))
}
