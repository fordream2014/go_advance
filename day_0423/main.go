package main

import (
	"context"
	"fmt"
	"go_advance/day_0423/dao"
)

func main() {
	ctx := context.Background()
	userDao := dao.NewUserDao()

	uid := "10000"
	userInfo, err := userDao.GetUserInfo(ctx, uid)
	if err != nil {
		fmt.Printf("error, %v\n", err)
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("user info:%v",userInfo)
}
