package main

import (
	"context"
	"fmt"
	"github.com/hankeyyh/mxshop_user_srv/model"
)

func main() {
	user := model.UserInstance()
	recList, err := user.BatchUser(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(recList)
	}
}
