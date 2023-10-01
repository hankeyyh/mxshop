package form

// PassWordLoginForm 注册请求
type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 图形验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

// 用户注册
type RegisterUserForm struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Nickname string `json:"nick_name" form:"nick_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=20"`
}
