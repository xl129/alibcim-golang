package im

import (
	"encoding/json"
	"strconv"
)

//群成员退出
type TribeQuit struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_quit_response"`
	}
}

func NewTribeQuit() *TribeQuit {
	return &TribeQuit{Paras: make(map[string]string)}
}

func (u *TribeQuit) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.quit"
}

func (u *TribeQuit) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeQuit) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeQuit) SetApiParas(json string) {
}

func (u *TribeQuit) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeQuit) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeQuit) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeQuit) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
