package api

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/client"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/log"
	middlewares "github.com/hankeyyh/mxshop/mxshop-api/user-web/middleware"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/proto"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/validators"
	"github.com/mojocn/base64Captcha"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"time"
)

var captchaStore = base64Captcha.DefaultMemStore

// PassWordLoginForm 注册请求
type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 图形验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

// HandleValidatorError 处理表单验证错误
func HandleValidatorError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	// 根据客户端语言获取翻译器
	lang := ctx.Request.Header.Get("Accept-Language")
	translator := validators.GetTranslator(lang)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeErrPrefix(errs.Translate(translator)),
	})
	return
}

func removeErrPrefix(e map[string]string) map[string]string {
	var res = make(map[string]string)
	for key, val := range e {
		//key = strings.SplitN(key, ".", 2)[1]
		key = key[strings.Index(key, ".")+1:]
		res[key] = val
	}
	return res
}

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

	// 图形验证码验证 测试环境暂时注释掉
	//if !captchaStore.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"captcha": "验证码错误",
	//	})
	//	return
	//}

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
