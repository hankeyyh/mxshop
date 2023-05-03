package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func NewMobileCVT() CustomValidateTranslation {
	return CustomValidateTranslation{
		Tag:          "mobile",
		ValidateFunc: validateMobile,
		CustomRegisFunc: func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		},
		CustomTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		},
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
