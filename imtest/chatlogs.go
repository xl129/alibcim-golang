package imtest

import (
	"ali/im"
	"fmt"
	"time"
)

func CGet() {
	u := im.NewChatLogsGet()
	apiuser1 := im.NewApiUser("Test4")
	apiuser2 := im.NewApiUser("Test5")
	u.SetApiUser1(apiuser1)
	u.SetApiUser2(apiuser2)
	u.SetCount(100)
	u.SetBegin(time.Now().Unix() - 86400*5)
	u.SetEnd(time.Now().Unix() + 86400)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
