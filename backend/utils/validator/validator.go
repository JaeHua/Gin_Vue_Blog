package validator

import (
	"backend/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// Validate 数据验证模块
func Validate(data any) (string, int) {
	// 创建一个新的验证器实例
	validate := validator.New()
	// 创建一个新的翻译器实例，指定中文翻译器
	uni := unTrans.New(zh_Hans_CN.New())
	// 注册默认的中文翻译
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	// 注册一个标签名称函数，用于获取结构体标签中的 `label` 字段
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	// 验证传入的数据结构体
	err = validate.Struct(data)
	if err != nil {
		// 如果验证失败，遍历错误并翻译为中文返回
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
