package base

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"
)

// 获取sdk 配置信息
func GetConfig() (sdkConfig *SdkConfig, err error) {
	_, filePath, _, _ := runtime.Caller(0)
	dirPath := path.Dir(filePath)
	file, err := ioutil.ReadFile(dirPath + "/../config/config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(file, &sdkConfig)
	return
}
