package im

import (
	"strconv"
	"encoding/json"
)

//获取群信息
type TribeGetInfo struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			Tribe            `json:"tribe_info"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_gettribeinfo_response"`
	}
}

func (u *TribeGetInfo) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.gettribeinfo"
}

func (u *TribeGetInfo) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id 错误"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeGetInfo) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeGetInfo) SetApiParas(json string) {
}

func (u *TribeGetInfo) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeGetInfo) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeGetInfo) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeGetInfo) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
