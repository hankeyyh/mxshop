package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/client"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/form"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/log"
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-03-01"))
	return []byte(stamp), nil
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
					"msg": e.Message(),
				})
			}
			return
		}
	}
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

// 注册用户
func Register(ctx *gin.Context) {
	registerForm := form.RegisterUserForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	userSvrClient := client.UserSvrClient
	rsp, err := userSvrClient.CreateUser(ctx, &proto.CreateUserInfo{
		Nickname: registerForm.Nickname,
		Password: registerForm.Password,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := UserResponse{
		Id:       rsp.GetId(),
		NickName: rsp.GetNickname(),
		Birthday: JsonTime(time.Unix(int64(rsp.GetBirthday()), 0)),
		Gender:   rsp.GetGender(),
		Mobile:   rsp.GetMobile(),
	}
	ctx.JSON(http.StatusOK, result)
}
