package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type MobileValidate struct {
	Tag          string
	v            *validator.Validate
	ValidateFunc validator.Func
}

func NewMobileCVT(v *validator.Validate) CustomValidator {
	return &MobileValidate{
		Tag:          "mobile",
		v:            v,
		ValidateFunc: validateMobile,
	}
}

func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}

func (m *MobileValidate) RegisterValidation() error {
	return m.v.RegisterValidation(m.Tag, m.ValidateFunc)
}

func (m *MobileValidate) RegisterTranslation(translator ut.Translator) error {
	locale := translator.Locale()
	var translation string
	switch locale {
	case "zh":
		translation = "{0} 非法的手机号码！"
	default:
		translation = "{0} invalid mobile!"
	}

	return m.v.RegisterTranslation(m.Tag, translator, func(ut ut.Translator) error {
		return ut.Add(m.Tag, translation, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(m.Tag, fe.Field())
		return t
	})
}
