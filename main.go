/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"gbv2/cmd"
	_ "gbv2/docs"
)

// @title blogV2 API文档
// @version 1.0
// @description blogV2 API文档
// @host 127.0.0.1:8087
// @BasePath /
func main() {
	cmd.Execute()
}
