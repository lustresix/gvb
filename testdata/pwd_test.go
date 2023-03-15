package testdata

import (
	"fmt"
	"gbv2/utils/pwd"
	"testing"
)

func TestPwd(t *testing.T) {
	fmt.Println(pwd.HashPwd("1234"))
}
