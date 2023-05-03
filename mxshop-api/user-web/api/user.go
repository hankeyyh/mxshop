package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/log"
	middlewares "mxshop-api/user-web/middleware"
	"mxshop-api/user-web/proto"
	"mxshop-api/user-web/validators"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-03-01"))
	return []byte(stamp), nil
}

// PassWordLoginForm 注册请求
type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 图形验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

// UserResponse 返回结构
type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	//Birthday string `json:"birthday"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}

// HandleGrpcErrorToHttp 将grpc的code转换成http的状态码
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "未找到内容",
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

// HandleValidatorError 处理表单验证错误
func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeErrPrefix(errs.Translate(validators.DefaultTranslator())),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	log.Info(ctx, "访问用户: ", log.Any("userId", userId))

	userSrvConf := config.DefaultConfig().UserSrvInfo
	addr := fmt.Sprintf("%s:%d", userSrvConf.Host, userSrvConf.Port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(context.Background(), "grpc.Dial fail", log.Any("err", err))
	}
	client := proto.NewUserClient(conn)

	pn := ctx.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		log.Error(ctx, "GetUserList fail", log.Any("err", err))
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	data := rsp.GetData()
	result := make([]interface{}, 0, len(data))
	for _, val := range data {
		result = append(result, &UserResponse{
			Id:       val.GetId(),
			NickName: val.GetNickname(),
			Birthday: JsonTime(time.Unix(int64(val.GetBirthday()), 0)),
			Gender:   val.GetGender(),
			Mobile:   val.GetMobile(),
		})
	}
	ctx.JSON(http.StatusOK, result)
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

	// 根据mobile查询用户
	userSrvConf := config.DefaultConfig().UserSrvInfo
	addr := fmt.Sprintf("%s:%d", userSrvConf.Host, userSrvConf.Port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(ctx, "grpc.Dial fail", log.Any("err", err))
	}
	client := proto.NewUserClient(conn)

	// 获取用户
	user, err := client.GetUserByMobile(context.Background(), &proto.MobileRequest{
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
	rsp, err := client.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
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
