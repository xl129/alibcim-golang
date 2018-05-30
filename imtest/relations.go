package imtest

import (
	"ali/im"
	"fmt"
)

func RGet() {
	u := &im.RelationsGet{Paras: make(map[string]string)}
	apiuser := im.NewApiUser("Test4")
	u.SetApiUser(apiuser)
	u.SetBegDate("20180528")
	u.SetEndDate("20180530")
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}