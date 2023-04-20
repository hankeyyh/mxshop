package test

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop-srv/user-srv/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"testing"
)

var (
	conn   *grpc.ClientConn
	client proto.UserClient
)

func init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = proto.NewUserClient(conn)
}

func TestGetUserList(t *testing.T) {
	defer conn.Close()
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestGetUserByMobile(t *testing.T) {
	defer conn.Close()
	rsp, err := client.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "657834",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestGetUserById(t *testing.T) {
	defer conn.Close()
	rsp, err := client.GetUserById(context.Background(), &proto.IdRequest{
		Id: 7,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestCreateUser(t *testing.T) {
	defer conn.Close()
	rsp, err := client.CreateUser(context.Background(), &proto.CreateUserInfo{
		Nickname: "q1",
		Password: "44444",
		Mobile:   "15514567",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestUpdateUser(t *testing.T) {
	defer conn.Close()
	rsp, err := client.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       2,
		Nickname: "yyh",
		Gender:   "female",
		Birthday: 18446744073709522816,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestCheckPassWord(t *testing.T) {
	defer conn.Close()
	pwd := []byte("44444789564")
	hpwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	rsp, err := client.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
		Password:          string(pwd),
		EncryptedPassword: string(hpwd),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetSuccess())
}
