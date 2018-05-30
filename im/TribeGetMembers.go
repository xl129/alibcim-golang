package im

import (
	"encoding/json"
	"strconv"
)

//群成员获取
type TribeGetMembers struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeUserList struct {
				TribeUser [] TribeUser `json:"tribe_user"`
			} `json:"tribe_user_list"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_getmembers_response"`
	}
}

func (u *TribeGetMembers) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.getmembers"
}

func (u *TribeGetMembers) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeGetMembers) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeGetMembers) SetApiParas(json string) {
}

func (u *TribeGetMembers) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeGetMembers) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeGetMembers) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeGetMembers) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
