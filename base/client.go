package base

// 返回客户端
func NewClient(c Credentials) (client *Client, err error) {
	// 读取配置
	sdkConfig, err := GetConfig()
	if err != nil {
		return
	}
	client = &Client{
		SdkConfig: SdkConfig{
			Host:                sdkConfig.Host,
			GetDrawSelector:     sdkConfig.GetDrawSelector,
			PutBatchTask:        sdkConfig.PutBatchTask,
			PutBatchControlTask: sdkConfig.PutBatchControlTask,
		},
		Credentials: Credentials{
			Apikey:    c.Apikey,
			Apisecret: c.Apisecret,
		},
	}

	return
}
