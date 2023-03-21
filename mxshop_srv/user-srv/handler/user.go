package handler

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop_user_srv/model"
	"github.com/hankeyyh/mxshop_user_srv/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	proto.UnimplementedUserServer
}

func (u UserService) GetUserList(ctx context.Context, request *proto.PageInfo) (*proto.UserListResonse, error) {
	pn := request.GetPn()
	psize := request.GetPSize()

	total, err := model.UserInstance().GetUserCnt()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userList, err := model.UserInstance().BatchUser(int(pn), int(psize))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var data = make([]*proto.UserInfoResponse, 0, len(userList))
	for _, user := range userList {
		data = append(data, &proto.UserInfoResponse{
			Id:       user.Id,
			Password: user.Password,
			Mobile:   user.Mobile,
			Nickname: user.Nickname,
			Birthday: user.Birthday,
			Gender:   user.Gender,
			Role:     user.Role,
		})
	}

	rsp := new(proto.UserListResonse)
	rsp.Total = int32(total)
	rsp.Data = data
	return rsp, nil
}

func (u UserService) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) CreateUser(ctx context.Context, info *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) UpdateUser(ctx context.Context, info *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) CheckPassWord(ctx context.Context, info *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}
