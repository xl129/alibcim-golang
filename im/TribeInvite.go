package im

import (
	"encoding/json"
	"strconv"
)

//群邀请加入
type TribeInvite struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Members []ApiUser
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_invite_response"`
	}
}

func NewTribeInvite() *TribeInvite {
	return &TribeInvite{Paras: make(map[string]string)}
}

func (u *TribeInvite) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.invite"
}

func (u *TribeInvite) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if len(u.Members) > 1000 {
		return false, "最多1000个群成员"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeInvite) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeInvite) SetApiParas(json string) {
}

func (u *TribeInvite) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeInvite) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeInvite) SetMemberList(list []ApiUser) {
	bytes, err := json.Marshal(list)
	if err != nil {
		return
	}
	u.Members = list
	u.SetApiOtherPara("members", string(bytes))
}

func (u *TribeInvite) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeInvite) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
