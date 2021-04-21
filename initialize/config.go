package initialize

import (
	"fmt"
	"github/qm012/nacos-adress/config"
	"github/qm012/nacos-adress/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func InitServerConfig() {
	file, err := os.Open("./application.yaml")
	if err != nil {
		panic(fmt.Sprintf("init config file err: %v", err.Error()))
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("read config file data err: %v", err.Error()))
	}
	var server = &config.Server{}
	err = yaml.Unmarshal(all, server)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal config file data err: %v", err.Error()))
	}
	global.Server = server
}
