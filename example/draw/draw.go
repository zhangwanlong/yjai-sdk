package main

import (
	"fmt"
	"github.com/zhangwanlong/yjai-sdk/base"
	"github.com/zhangwanlong/yjai-sdk/service/draw"
)

var client *draw.Client

func init() {
	// 初始化获取sdk 客户端
	Credentials := base.Credentials{
		Apikey:    "dcbc12e70664bfcff9f86785326245ee",
		Apisecret: "ddc55987fd3c4dc997159a32c7162005",
	}
	var err error
	client, err = draw.NewClient(Credentials)
	if err != nil {
		return
	}
}

func main() {
	// 获取绘画风格
	GetDrawSelector()
	//批量绘画任务
	//PutBatchTask()
}

// 批量绘画任务
func PutBatchTask() {
	putBatchTask := base.PutBatchTask{
		Online:        0,
		TaskNum:       66,
		ResLevel:      0,
		EnableMix:     0,
		Prompt:        "太阳",
		Ratio:         1,
		UseModel336:   1,
		InitStrength:  50,
		Engine:        "stable_diffusion",
		GuidenceScale: 7.5,
		IsVipChannel:  0,
		CallbackType:  "end",
		CallbackUrl:   "http://47.111.11.42:8082/painting-open-api/site/callback_task",
	}
	ret, err := client.PutBatchTask(putBatchTask)
	// 打印数据
	fmt.Println(ret)
	//打印数据
	fmt.Println(err)
}

// 获取绘画风格
func GetDrawSelector() {
	ret, err := client.GetDrawSelector()
	// 打印数据
	fmt.Println(ret)
	//打印数据
	fmt.Println(err)
}
