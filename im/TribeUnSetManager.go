package im

import (
	"encoding/json"
	"strconv"
)

//群取消管理员
type TribeUnSetManager struct {
	User   ApiUser `json:"user"`
	Tid    uint64  `json:"tribe_id"`
	Member ApiUser `json:"member"`
	Paras  map[string]string
	ResponseData struct {
		Data struct {
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_unsetmanager_response"`
	}
}

func NewTribeUnSetManager() *TribeUnSetManager {
	return &TribeUnSetManager{Paras: make(map[string]string)}
}

func (u *TribeUnSetManager) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.unsetmanager"
}

func (u *TribeUnSetManager) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.Member.Uid == "" {
		return false, "Member.uid错误"
	}
	if u.Tid <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeUnSetManager) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeUnSetManager) SetApiParas(json string) {
}

func (u *TribeUnSetManager) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeUnSetManager) SetMember(member ApiUser) () {
	u.Member = member
	bytes, err := json.Marshal(member)
	if err != nil {
		return
	}
	u.SetApiOtherPara("member", string(bytes))
}

func (u *TribeUnSetManager) SetTribeId(tribe_id uint64) () {
	u.Tid = tribe_id
	u.SetApiOtherPara("tid", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeUnSetManager) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeUnSetManager) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
