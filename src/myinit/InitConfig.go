package myinit

import (
	"fmt"
	"github.com/olebedev/config"
	"os"
)

var GCF *config.Config //global config

func InitConf() {
	pwd, _ := os.Getwd()
	configPath := pwd + "/configs/application.yml"
	fmt.Println(configPath)
	GCF, _ = config.ParseYamlFile(configPath)
}
