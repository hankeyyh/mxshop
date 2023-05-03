package validators

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// CustomValidateTranslation 自定义tag的验证&翻译
type CustomValidateTranslation struct {
	Tag             string
	ValidateFunc    validator.Func
	CustomRegisFunc validator.RegisterTranslationsFunc
	CustomTransFunc validator.TranslationFunc
}

var cvtList []CustomValidateTranslation

var defaultTranslator ut.Translator

func DefaultTranslator() ut.Translator {
	return defaultTranslator
}

func Init(locale string) error {
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

	// 默认翻译
	uni := ut.New(en.New(), zh.New(), en.New())
	defaultTranslator, ok = uni.GetTranslator(locale)
	if !ok {
		return errors.New("locale translator not found")
	}

	var err error
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, defaultTranslator)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, defaultTranslator)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, defaultTranslator)
	}
	if err != nil {
		return errors.New("RegisterDefaultTranslations fail")
	}

	// 自定义验证&翻译
	cvtList = append(cvtList, NewMobileCVT())

	for _, vt := range cvtList {
		if err = v.RegisterValidation(vt.Tag, vt.ValidateFunc); err != nil {
			return errors.New(fmt.Sprintf("[%s]reg validate fail, %s", vt.Tag, err.Error()))
		}
		if err = v.RegisterTranslation(vt.Tag, defaultTranslator, vt.CustomRegisFunc, vt.CustomTransFunc); err != nil {
			return errors.New(fmt.Sprintf("[%s]reg trans fail, %s", vt.Tag, err.Error()))
		}
	}

	return nil
}
