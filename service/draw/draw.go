package draw

import (
	"github.com/google/go-querystring/query"
	"github.com/zhangwanlong/yjai-sdk/base"
	"net/url"
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

// 获取绘画风格
func (c *Client) GetDrawSelector() (ret string, err error) {
	requestUrl := c.SdkConfig.Host + c.SdkConfig.GetDrawSelector
	httpClient := base.NewHttpClient(c.Credentials.Apikey, c.Credentials.Apisecret, 10)
	queryValues := url.Values{}
	ret, err = httpClient.RequestPost(requestUrl, queryValues)
	return
}

// 开启批量绘画任务
func (c *Client) PutBatchTask(data base.PutBatchTask) (ret string, err error) {
	//验证参数
	err = base.ValidatorData(data)
	if err != nil {
		return
	}
	queryValues, err := query.Values(data)

	if err != nil {
		return
	}
	requestUrl := c.SdkConfig.Host + c.SdkConfig.PutBatchTask
	httpClient := base.NewHttpClient(c.Credentials.Apikey, c.Credentials.Apisecret, 10)
	ret, err = httpClient.RequestPost(requestUrl, queryValues)
	return
}

// 开启批量绘画任务
func (c *Client) PutBatchControlTask(data base.PutBatchControlTask) (ret string, err error) {
	//验证参数
	err = base.ValidatorData(data)
	if err != nil {
		return
	}
	queryValues, err := query.Values(data)

	if err != nil {
		return
	}
	requestUrl := c.SdkConfig.Host + c.SdkConfig.PutBatchControlTask
	httpClient := base.NewHttpClient(c.Credentials.Apikey, c.Credentials.Apisecret, 10)
	ret, err = httpClient.RequestPost(requestUrl, queryValues)
	return
}
