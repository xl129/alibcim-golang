package im

import (
	"encoding/json"
	"strconv"
)

//创建群
type TribeCreate struct {
	User      ApiUser `json:"user"`
	TribeName string  `json:"tribe_name"`
	Notice    string  `json:"notice"`
	TribeType int     `json:"tribe_type"`
	Members   []ApiUser
	Paras     map[string]string
	ResponseData struct {
		Data struct {
			Tribe            `json:"tribe_info"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_create_response"`
	}
}

func (u *TribeCreate) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.create"
}

func (u *TribeCreate) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeName == "" {
		return  false, "缺少参数TribeName"
	}
	if u.Notice == "" {
		return  false, "缺少参数Notice"
	}
	if len(u.Members) > 1000 {
		return false, "最多1000个群成员"
	}
	return true, ""
}

func (u *TribeCreate) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeCreate) SetApiParas(json string) {
}

func (u *TribeCreate) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeCreate) SetTribeName(tribe_name string) () {
	u.TribeName = tribe_name
	u.SetApiOtherPara("tribe_name", tribe_name)
}

func (u *TribeCreate) SetNotice(notice string) () {
	u.Notice = notice
	u.SetApiOtherPara("notice", notice)
}

func (u *TribeCreate) SetTribeType(tribe_type int) () {
	u.TribeType = tribe_type
	u.SetApiOtherPara("tribe_type", strconv.Itoa(tribe_type))
}

func (u *TribeCreate) SetMemberList(list []ApiUser) {
	bytes, err := json.Marshal(list)
	if err != nil {
		return
	}
	u.Members = list
	u.SetApiOtherPara("members", string(bytes))
}

func (u *TribeCreate) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeCreate) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
