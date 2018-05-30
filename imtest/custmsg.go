package imtest

import (
	"ali/im"
	"fmt"
)

func MPush() {
	u := &im.CustMsgPush{Paras: make(map[string]string)}
	u.SetFromUser("Test4")
	var uids = make([]string,3)
	uids[0]="Test5"
	uids[1]="Test6"
	uids[2]="Test0"
	u.SetToUsers(uids)
	u.SetSummary("测试一下")
	u.SetData("测试")
	u.SetAps("ios apns push")
	u.SetApnsParam("apns推送的附带数据")
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
