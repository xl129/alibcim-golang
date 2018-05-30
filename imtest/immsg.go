package imtest

import (
	"ali/im"
	"fmt"
)

func IPush() {
	u := &im.IMMsgPush{Paras: make(map[string]string)}
	u.SetFromUser("Test4")
	var uids = make([]string, 3)
	uids[0] = "Test5"
	uids[1] = "Test6"
	uids[2] = "Test0"
	u.SetToUsers(uids)
	u.SetContext("测试一下")
	u.SetMsgType(0)
	u.SetMediaAttr("")
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
