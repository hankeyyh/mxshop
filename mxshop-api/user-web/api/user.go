package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/client"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/log"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/proto"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	Id       int32    `json:"id"`
	NickName string   `json:"name"`
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

func removeErrPrefix(e map[string]string) map[string]string {
	var res = make(map[string]string)
	for key, val := range e {
		//key = strings.SplitN(key, ".", 2)[1]
		key = key[strings.Index(key, ".")+1:]
		res[key] = val
	}
	return res
}

func GetUserList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	log.Info(ctx, "访问用户: ", log.Any("userId", userId))

	pn := ctx.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	userSvrClient := client.UserSvrClient
	rsp, err := userSvrClient.GetUserList(context.Background(), &proto.PageInfo{
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
