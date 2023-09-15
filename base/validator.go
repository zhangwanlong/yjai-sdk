package base

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

func ValidatorData(data interface{}) (err error) {
	// 中文翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	//实例化验证器
	validate := validator.New()
	// 注册翻译器到校验器
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}

	errs := validate.Struct(data)
	var errStr []string
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errStr = append(errStr, err.Translate(trans))
		}
	}
	if len(errStr) > 0 {
		err = errors.New(strings.Join(errStr, ";"))
	}
	return
}
