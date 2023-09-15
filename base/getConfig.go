package base

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

// 获取sdk 配置信息
func GetConfig() (sdkConfig *SdkConfig, err error) {
	file, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(file, &sdkConfig)
	return
}
