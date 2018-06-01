package im

import (
	"encoding/json"
	"strconv"
)

//群踢出成员
type TribeExpel struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Member  ApiUser `json:"member"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_expel_response"`
	}
}

func NewTribeExpel() *TribeExpel {
	return &TribeExpel{Paras: make(map[string]string)}
}

func (u *TribeExpel) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.expel"
}

func (u *TribeExpel) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.Member.Uid == "" {
		return false, "Member.uid错误"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeExpel) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeExpel) SetApiParas(json string) {
}

func (u *TribeExpel) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeExpel) SetMember(member ApiUser) () {
	u.Member = member
	bytes, err := json.Marshal(member)
	if err != nil {
		return
	}
	u.SetApiOtherPara("member", string(bytes))
}

func (u *TribeExpel) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeExpel) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeExpel) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
