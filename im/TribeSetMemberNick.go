package im

import (
	"encoding/json"
	"strconv"
)

//设置群成员昵称
type TribeSetMemberNick struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Member  ApiUser `json:"member"`
	Nick    string  `json:"nick"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_setmembernick_response"`
	}
}

func NewTribeSetMemberNick() *TribeSetMemberNick {
	return &TribeSetMemberNick{Paras: make(map[string]string)}
}

func (u *TribeSetMemberNick) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.setmembernick"
}

func (u *TribeSetMemberNick) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.Member.Uid == "" {
		return false, "Member.uid错误"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	if u.Nick == "" {
		return false, "Nick错误"
	}
	return true, ""
}

func (u *TribeSetMemberNick) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeSetMemberNick) SetApiParas(json string) {
}

func (u *TribeSetMemberNick) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeSetMemberNick) SetMember(member ApiUser) () {
	u.Member = member
	bytes, err := json.Marshal(member)
	if err != nil {
		return
	}
	u.SetApiOtherPara("member", string(bytes))
}

func (u *TribeSetMemberNick) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeSetMemberNick) SetNick(nick string) () {
	u.Nick = nick
	u.SetApiOtherPara("nick", nick)
}

func (u *TribeSetMemberNick) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeSetMemberNick) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
