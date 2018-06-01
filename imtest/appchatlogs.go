package imtest

import (
	"ali/im"
	"time"
	"fmt"
)

func AGet() {
	u := im.NewAppChatLogsGet()
	u.SetCount(100)
	u.SetBeg(time.Now().Unix() - 86400*5)
	u.SetEnd(time.Now().Unix() + 86400)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
