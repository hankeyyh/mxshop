package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"mxshop-api/user-web/log"
	"net/http"
)

var captchaStore = base64Captcha.DefaultMemStore

// GetCaptcha 获取图形验证码
func GetCaptcha(ctx *gin.Context) {
	// 生成图形验证码
	var driver = base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		log.Error(ctx, "生成验证码错误,: ", log.Any("err", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证码生成失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
