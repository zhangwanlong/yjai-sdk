package base

type Client struct {
	Credentials Credentials `json:"credentials,omitempty"`
	SdkConfig   SdkConfig   `json:"sdk_Config,omitempty"`
}

// 凭证信息
type Credentials struct {
	Apikey    string `json:"apikey,omitempty"`
	Apisecret string `json:"apisecret,omitempty"`
}

// sdk 配置信息
type SdkConfig struct {
	Host                string `yaml:"host"`
	GetDrawSelector     string `yaml:"getDrawSelector"`
	PutBatchTask        string `yaml:"putBatchTask"`
	PutBatchControlTask string `yaml:"putBatchControlTask"`
}

// 批量绘画任务
type PutBatchTask struct {
	Online            int64   `json:"online,omitempty" url:"online"`                                         //是否在线体验，可以不传此字段，默认 0
	TaskNum           int64   `json:"task_num,omitempty" url:"task_num" validate:"min=1,max=4"`              //批量绘制的任务数：task_num取值必须是1, 2, 4之中的值
	ResLevel          int64   `json:"res_level,omitempty" url:"res_level"`                                   //精绘模式：不需要可以不传此字段，默认 0
	EnableMix         int64   `json:"enable_mix,omitempty" url:"enable_mix"`                                 //是否开启混合，不需要可以不传此字段，默认 0
	Prompt            string  `json:"prompt,omitempty" validate:"required" url:"prompt"`                     //描述词，不能为空
	Ratio             int64   `json:"ratio,omitempty" validate:"min=0,max=4" url:"ratio"`                    //尺寸比，0(4:3) 1(3:4) 2(1:1) 3（16:9）4:(9:16)
	Style             string  `json:"style,omitempty" url:"style"`                                           //风格，不需要可以不传此字段，默认为空
	SubStyle          string  `json:"sub_style,omitempty" url:"sub_style"`                                   //子风格，不需要可以不传此字段，默认为空
	StepsMode         int64   `json:"steps_mode,omitempty" url:"steps_mode"`                                 //步长参数，不需要可以不传此字段，默认为 0；取值范围：0常规步长|1:加长步长
	UseModel336       int64   `json:"use_model_336,omitempty" url:"use_model_336"`                           //是否使用 336 模型，不需要可以不传此字段，默认 1
	EnableFaceEnhance int64   `json:"enable_face_enhance,omitempty" url:"enable_face_enhance"`               //是否开启面部增强。不需要可以不传此字段，默认 0
	IsLastLayerSkip   int64   `json:"is_last_layer_skip,omitempty" url:"is_last_layer_skip"`                 //是否开启色彩狂化，不需要可以不传此字段，默认 0
	InitImage         string  `json:"init_image,omitempty" url:"init_image"`                                 //参考图，不需要可以不传此字段，http图片和https图片均可。上传图片接口需要自己搞定，图片分辨率不要超过1000px*1000px，体积尽量小。不能传null!!!
	InitStrength      float32 `json:"init_strength,omitempty" url:"init_strength"`                           //图片相关性，0-70，不需要可以不传此字段，默认 50
	Engine            string  `json:"engine,omitempty" url:"engine"`                                         //引擎值，默认为空【需要说明如何取值】
	GuidenceScale     float32 `json:"guidence_scale,omitempty" validate:"min=3,max=25" url:"guidence_scale"` //引导力，默认为7.5，取值范围：3-25
	IsVipChannel      int64   `json:"is_vip_channel,omitempty" url:"is_vip_channel"`                         //是否是用VIP通道，不需要可以不传此字段，默认 0
	CallbackUrl       string  `json:"callback_url,omitempty" url:"callback_url"`                             //回调接口，字段可以不传。如果填了，会以回调的方式同步当前任务的进度
	CallbackType      string  `json:"callback_type,omitempty" url:"callback_type"`                           //progress和end,end任务结束的时候(失败也会)回调,progress有进度（失败也会）就会回调
}

// 批量图生图
type PutBatchControlTask struct {
	PutBatchTask
	ControlImage     string  `json:"control_image,omitempty" url:"control_image"`         //参考图url
	ControlProcessor string  `json:"control_processor,omitempty" url:"control_processor"` //轮廓类型 lineart_anime:动漫物体  lineart：真实物体 none:默认无选择
	ControlModel     string  `json:"control_model,omitempty" url:"control_model"`
	GuidanceStart    float32 `json:"guidance_start,omitempty" url:"guidance_start"`
	GuidanceEnd      float32 `json:"guidance_end,omitempty" url:"guidance_end"`
	ThresholdA       float32 `json:"threshold_a,omitempty" url:"threshold_a"`
	ThresholdB       float32 `json:"threshold_b,omitempty" url:"threshold_b"`
	Weight           float32 `json:"weight,omitempty" url:"weight"`
	ControlMode      int64   `json:"control_mode,omitempty" url:"control_mode"`
}
