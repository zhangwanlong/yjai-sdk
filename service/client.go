package service

import (
	"github.com/zhangwanlong/yjai-sdk/base"
)

type Client struct {
	Credentials base.Credentials `json:"credentials,omitempty"`
	SdkConfig   base.SdkConfig   `json:"sdkConfig,omitempty"`
}

// 返回客户端
func NewClient(c base.Credentials) (client *Client, err error) {
	// 读取配置
	sdkConfig, err := base.GetConfig()
	if err != nil {
		return
	}
	client = &Client{
		SdkConfig: base.SdkConfig{
			Host:                sdkConfig.Host,
			GetDrawSelector:     sdkConfig.GetDrawSelector,
			PutBatchTask:        sdkConfig.PutBatchTask,
			PutBatchControlTask: sdkConfig.PutBatchControlTask,
		},
		Credentials: base.Credentials{
			Apikey:    c.Apikey,
			Apisecret: c.Apisecret,
		},
	}

	return
}
