package imtest

import (
	"ali/im"
	"fmt"
	"time"
)

func TAdd() {
	apiuser := im.NewApiUser("Test4")
	var u = &im.TribeCreate{Paras: make(map[string]string)}
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
	var u = &im.TribeGetInfo{Paras: make(map[string]string)}
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
	var u = &im.TribeQuit{Paras: make(map[string]string)}
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
	var u = &im.TribeJoin{Paras: make(map[string]string)}
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
	var u = &im.TribeExpel{Paras: make(map[string]string)}
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
	var u = &im.TribeSetManager{Paras: make(map[string]string)}
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
	var u = &im.TribeUnSetManager{Paras: make(map[string]string)}
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
	var u = &im.TribeDismiss{Paras: make(map[string]string)}
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
	var u = &im.TribeInvite{Paras: make(map[string]string)}
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
	var u = &im.TribeGetMembers{Paras: make(map[string]string)}
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
	var u = &im.TribeGetAllTribes{Paras: make(map[string]string)}
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
	var u = &im.TribeModifyTribeInfo{Paras: make(map[string]string)}
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
	var u = &im.TribeSetMemberNick{Paras: make(map[string]string)}
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
	var u = &im.TribeSendMsg{Paras: make(map[string]string)}
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
	u := &im.TribeLogsGet{Paras: make(map[string]string)}
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
