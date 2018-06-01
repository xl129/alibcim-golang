package imtest

import (
	"ali/im"
	"strconv"
	"fmt"
)

func UAdd() {
	var UserList = make([]im.User, 1)
	i := 6
	item := im.User{
		UserId:   "Test" + strconv.Itoa(i),
		Password: "Pass" + strconv.Itoa(i),
	}
	UserList[0] = item
	var u = im.NewUserAdd()
	u.SetUserList(UserList)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func UDelete() {
	// Test1 Test2 已经删除
	var UserIDList = make([]string, 2)
	UserIDList[0] = "Test1"
	UserIDList[1] = "Test2"
	var u = im.NewUserDelete()
	u.SetUserIDsList(UserIDList)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func UUpdate() {
	var UserList = make([]im.User, 1)
	i := 0
	item := im.User{
		UserId:   "Test" + strconv.Itoa(i),
		Password: "Pass" + strconv.Itoa(i),
	}
	UserList[0] = item
	var u = im.NewUserUpdate()
	u.SetUserList(UserList)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func UGetInfo() {
	var UserIDList = make([]string, 7)
	for i := 0; i < 7; i++ {
		UserIDList[i] = "Test" + strconv.Itoa(i)
	}
	var u = im.NewUserGetInfo()
	u.SetUserIDsList(UserIDList)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
