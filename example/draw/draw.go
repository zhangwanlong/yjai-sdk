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
	//ret, err := GetDrawSelector()
	//批量绘画任务
	//ret, err := PutBatchTask()
	// 打印数据
	ret, err := PutBatchControlTask()

	fmt.Println(ret)
	//打印数据
	fmt.Println(err)
}

// 获取绘画风格
func GetDrawSelector() (ret string, err error) {
	ret, err = client.GetDrawSelector()
	return
}

// 批量绘画任务
func PutBatchTask() (ret string, err error) {
	putBatchTask := base.PutBatchTask{
		Online:        0,
		TaskNum:       2,
		ResLevel:      0,
		EnableMix:     0,
		Prompt:        "雨伞",
		Ratio:         1,
		UseModel336:   1,
		InitStrength:  50,
		Engine:        "stable_diffusion",
		GuidenceScale: 7.5,
		IsVipChannel:  0,
		CallbackType:  "end",
		CallbackUrl:   "http://47.111.11.42:8082/painting-open-api/site/callback_task",
	}
	ret, err = client.PutBatchTask(putBatchTask)
	return
}

// 批量图生图绘画任务
func PutBatchControlTask() (ret string, err error) {
	putBatchControlTask := base.PutBatchControlTask{
		PutBatchTask: base.PutBatchTask{
			Online:        0,
			TaskNum:       2,
			ResLevel:      0,
			EnableMix:     0,
			Prompt:        "卡通人物",
			Ratio:         1,
			UseModel336:   1,
			InitStrength:  50,
			Engine:        "mid_stable_diffusion",
			GuidenceScale: 7.5,
			IsVipChannel:  0,
		},
		ControlImage:     "https://app.yjai.art:30108/yijian-new-painting4/upload_user_image/13524263213/4c9846c2-60a7-4288-9249-21d957148c92.jpg",
		ControlProcessor: "lineart",
		ControlModel:     "control_v11p_sd15_lineart.pth",
		GuidanceStart:    0,
		GuidanceEnd:      1,
		ThresholdA:       0,
		ThresholdB:       0,
		Weight:           1,
		ControlMode:      0,
	}
	ret, err = client.PutBatchControlTask(putBatchControlTask)
	return
}
