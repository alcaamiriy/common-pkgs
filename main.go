package main

import (
	"github.com/alcaamiriy/common-pkgs/utils"
)

func main() {
	err := utils.LoadConfig(".")
	utils.InitLogger()
	if err != nil {
		utils.Logger.Printf("cannot load config: %v", err)
	}
}
