package validators

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// 自定义验证器接口
type CustomValidator interface {
	RegisterValidation() error
	RegisterTranslation(translator ut.Translator) error
}

var cvtList []CustomValidator

var translatorMap = make(map[string]ut.Translator)

func GetTranslator(locale string) ut.Translator {
	if translator, ok := translatorMap[locale]; ok {
		return translator
	}
	// 默认英文
	return translatorMap["en"]
}

func Init() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("get validate fail")
	}
	// tag name
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	var err error

	// 翻译
	uni := ut.New(en.New(), zh.New(), en.New())

	zhTranslator, _ := uni.GetTranslator("zh")
	enTranslator, _ := uni.GetTranslator("en")
	translatorMap["zh"] = zhTranslator
	translatorMap["en"] = enTranslator

	// 自定义验证
	cvtList = append(cvtList, NewMobileCVT(v))

	for _, cvt := range cvtList {
		// 注册验证器
		if err = cvt.RegisterValidation(); err != nil {
			return errors.New(fmt.Sprintf("reg validate fail, %s", err.Error()))
		}
		// 注册验证器中文翻译
		if err = cvt.RegisterTranslation(zhTranslator); err != nil {
			return errors.New(fmt.Sprintf("reg zh-trans fail, %s", err.Error()))
		}
		// 注册验证器英文翻译
		if err = cvt.RegisterTranslation(enTranslator); err != nil {
			return errors.New(fmt.Sprintf("reg en-trans fail, %s", err.Error()))
		}
	}

	return nil
}
