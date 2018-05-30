package im

import (
	"encoding/json"
	"strconv"
)

//群信息修改
type TribeModifyTribeInfo struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	TribeName string  `json:"tribe_name"`
	Notice    string  `json:"notice"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_modifytribeinfo_response"`
	}
}

func (u *TribeModifyTribeInfo) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.modifytribeinfo"
}

func (u *TribeModifyTribeInfo) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeName == "" {
		return  false, "缺少参数TribeName"
	}
	if u.Notice == "" {
		return  false, "缺少参数Notice"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeModifyTribeInfo) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeModifyTribeInfo) SetApiParas(json string) {
}

func (u *TribeModifyTribeInfo) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeModifyTribeInfo) SetTribeName(tribe_name string) () {
	u.TribeName = tribe_name
	u.SetApiOtherPara("tribe_name", tribe_name)
}

func (u *TribeModifyTribeInfo) SetNotice(notice string) () {
	u.Notice = notice
	u.SetApiOtherPara("notice", notice)
}

func (u *TribeModifyTribeInfo) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeModifyTribeInfo) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeModifyTribeInfo) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
