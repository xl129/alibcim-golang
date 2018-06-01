package imtest

import (
	"ali/im"
	"fmt"
	"time"
)

func TCreate() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeCreate()
	u.SetApiUser(apiuser)
	u.SetTribeName("Test测试一下")
	u.SetNotice("这是一个群公告")
	u.SetTribeType(0)
	apiuserlist := make([]im.ApiUser, 2)
	apiuserlist[0] = im.NewApiUser("Test0")
	apiuserlist[1] = im.NewApiUser("Test1")
	u.SetMemberList(apiuserlist)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TGetInfo() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeGetInfo()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TQuit() {
	apiuser := im.NewApiUser("Test0")
	var u = im.NewTribeQuit()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TJoin() {
	apiuser := im.NewApiUser("Test5")
	var u = im.NewTribeJoin()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TExpel() {
	apiuser := im.NewApiUser("Test4")
	apimember := im.NewApiUser("Test0")
	var u = im.NewTribeExpel()
	u.SetApiUser(apiuser)
	u.SetMember(apimember)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TSetManager() {
	apiuser := im.NewApiUser("Test4")
	apimember := im.NewApiUser("Test5")
	var u = im.NewTribeSetManager()
	u.SetApiUser(apiuser)
	u.SetMember(apimember)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TUnSetManager() {
	apiuser := im.NewApiUser("Test4")
	apimember := im.NewApiUser("Test5")
	var u = im.NewTribeUnSetManager()
	u.SetApiUser(apiuser)
	u.SetMember(apimember)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TDismiss() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeDismiss()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610946040)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TInvite() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeInvite()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	apiuserlist := make([]im.ApiUser, 1)
	apiuserlist[0] = im.NewApiUser("Test6")
	u.SetMemberList(apiuserlist)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TGetMembers() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeGetMembers()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TGetAllTribes() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeGetAllTribes()
	u.SetApiUser(apiuser)
	u.SetTribeTypes([]string{"0", "1"})
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TModifyTribeInfo() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeModifyTribeInfo()
	u.SetApiUser(apiuser)
	u.SetTribeName("改个名字")
	u.SetNotice("改个公告")
	u.SetTribeId(2610653485)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TSetMemberNick() {
	apiuser := im.NewApiUser("Test4")
	apimember := im.NewApiUser("Test4")
	var u = im.NewTribeSetMemberNick()
	u.SetApiUser(apiuser)
	u.SetMember(apimember)
	u.SetTribeId(2610653485)
	u.SetNick("测试一下")
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TSendMsg() {
	apiuser := im.NewApiUser("Test4")
	var u = im.NewTribeSendMsg()
	u.SetApiUser(apiuser)
	u.SetTribeId(2610653485)
	u.SetMsgMsgContent("测试一下")
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}

func TLogsGet() {
	u := im.NewTribeLogsGet()
	u.SetTribeId(2610653485)
	u.SetCount(100)
	u.SetBegin(time.Now().Unix() - 86400)
	u.SetEnd(time.Now().Unix() + 86400)
	ResponseData := im.Execute(u)
	if ResponseData.Code == 0 {
		fmt.Println(u.ResponseData)
		return
	}
	fmt.Println(ResponseData)
}
