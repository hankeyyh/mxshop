package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	middlewares "mxshop-api/user-web/middleware"
	"mxshop-api/user-web/model"
	"mxshop-api/user-web/proto"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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
		"error": removeErrPrefix(errs.Translate(global.Trans)),
	})
	return
}

func GetUserList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	zap.S().Info("访问用户: ", userId)

	userSrvConf := global.ServerConfig.UserSrvInfo
	addr := fmt.Sprintf("%s:%d", userSrvConf.Host, userSrvConf.Port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Panic("grpc.Dial fail, err: ", err.Error())
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
		zap.S().Error("GetUserList fail, err: ", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	data := rsp.GetData()
	result := make([]interface{}, 0, len(data))
	for _, val := range data {
		result = append(result, &response.UserResponse{
			Id:       val.GetId(),
			NickName: val.GetNickname(),
			Birthday: response.JsonTime(time.Unix(int64(val.GetBirthday()), 0)),
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
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	// 图形验证码验证
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	// 根据mobile查询用户
	userSrvConf := global.ServerConfig.UserSrvInfo
	addr := fmt.Sprintf("%s:%d", userSrvConf.Host, userSrvConf.Port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		zap.S().Panic("grpc.Dial fail, err: ", err.Error())
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
	token, err := j.CreateToken(model.CustomClaims{
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
