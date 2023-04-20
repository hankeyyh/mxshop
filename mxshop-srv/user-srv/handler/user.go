package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/hankeyyh/mxshop-srv/user-srv/model"
	"github.com/hankeyyh/mxshop-srv/user-srv/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	proto.UnimplementedUserServer
}

func (u UserService) convertToUserInfo(rec model.User) *proto.UserInfoResponse {
	var user = &proto.UserInfoResponse{
		Id:       rec.Id,
		Password: rec.Password,
		Mobile:   rec.Mobile,
		Nickname: rec.Nickname,
		Birthday: uint64(rec.Birthday.Unix()),
		Gender:   rec.Gender,
		Role:     rec.Role,
	}
	return user
}

func (u UserService) GetUserList(ctx context.Context, request *proto.PageInfo) (*proto.UserListResonse, error) {
	pn := request.GetPn()
	psize := request.GetPSize()

	total, err := model.UserInstance().GetUserCnt()
	if err != nil {
		return nil, err
	}

	userList, err := model.UserInstance().BatchUser(int(pn), int(psize))
	if err != nil {
		return nil, err
	}

	var data = make([]*proto.UserInfoResponse, 0, len(userList))
	for _, user := range userList {
		data = append(data, u.convertToUserInfo(user))
	}

	rsp := new(proto.UserListResonse)
	rsp.Total = int32(total)
	rsp.Data = data
	return rsp, nil
}

func (u UserService) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	mobile := request.GetMobile()
	rec, err := model.UserInstance().GetUserByMobile(mobile)
	if err != nil {
		return nil, err
	}
	return u.convertToUserInfo(rec), nil
}

func (u UserService) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	id := request.GetId()
	rec, err := model.UserInstance().GetUser(id)
	if err != nil {
		return nil, err
	}
	return u.convertToUserInfo(rec), nil
}

func (u UserService) CreateUser(ctx context.Context, request *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	mobile := request.GetMobile()
	nickName := request.GetNickname()
	password := request.GetPassword()

	_, err := model.UserInstance().GetUserByMobile(mobile)
	if err == nil {
		return nil, fmt.Errorf("用户已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// db异常
		return nil, err
	}

	hpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("密码加密错误")
	}
	var user = model.User{
		Mobile:   mobile,
		Nickname: nickName,
		Password: string(hpwd),
	}
	if err = model.UserInstance().CreateUser(&user); err != nil {
		return nil, err
	}

	return u.convertToUserInfo(user), err
}

func (u UserService) UpdateUser(ctx context.Context, request *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	id := request.GetId()
	gender := request.GetGender()
	nickName := request.GetNickname()
	birthday := request.GetBirthday()

	user, err := model.UserInstance().GetUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}
	b := int64(birthday)
	user.Gender = gender
	user.Nickname = nickName
	user.Birthday = time.Unix(b, 0).UTC()
	if err = model.UserInstance().UpdateUser(user); err != nil {
		return nil, err
	}
	rsp := new(emptypb.Empty)
	return rsp, nil
}

func (u UserService) CheckPassWord(ctx context.Context, request *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	pwd := request.GetPassword()
	encryptedPwd := request.GetEncryptedPassword()
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(pwd))

	rsp := new(proto.CheckResponse)
	rsp.Success = err == nil
	return rsp, nil
}
