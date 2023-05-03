package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/client"
	"mxshop-api/user-web/log"
	middlewares "mxshop-api/user-web/middleware"
	"mxshop-api/user-web/proto"
	"net/http"
	"time"
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

// PasswordLogin 密码登录
func PasswordLogin(ctx *gin.Context) {
	// 表单验证
	passwordLoginForm := PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	// 图形验证码验证
	if !captchaStore.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	userSvrClient := client.UserSvrClient

	// 获取用户
	user, err := userSvrClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "服务异常，登录失败",
				})
			}
		}
		return
	}

	// 校验密码
	rsp, err := userSvrClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
		Password:          passwordLoginForm.PassWord,
		EncryptedPassword: user.GetPassword(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"password": "服务异常，登录失败",
		})
		return
	}
	if !rsp.GetSuccess() {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"msg": "密码错误",
		})
		return
	}

	// 生成token
	j := middlewares.NewJWT()
	token, err := j.CreateToken(middlewares.CustomClaims{
		ID:          uint(user.GetId()),
		NickName:    user.GetNickname(),
		AuthorityId: uint(user.GetRole()),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + (24 * 3600 * 30), //30天过期
			Issuer:    "mxshop",
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.GetId(),
		"nick_name":  user.GetNickname(),
		"token":      token,
		"expired_at": (time.Now().Unix() + 3600*24*30) * 1000,
	})
}
