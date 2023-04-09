package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"mxshop-api/user-web/global"
	"reflect"
	"strings"
)

func InitTrans(locale string) (err error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	// 注册一个json tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New()
	enT := en.New()

	uni := ut.New(enT, zhT, enT)

	global.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
	}
	return nil
}
