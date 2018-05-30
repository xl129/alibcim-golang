package im

import (
	"encoding/json"
	"strings"
)

//批量获取用户信息
type UserGetInfo struct {
	UserIDList []string
	Paras      map[string]string
	ResponseData struct {
		Data struct {
			Userinfos struct {
				Userinfos []User `json:"userinfos"`
			} `json:"userinfos"`
		} `json:"openim_users_get_response"`
	}
}

func (u *UserGetInfo) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.users.get"
}

func (u *UserGetInfo) CheckApiParas() (bool, string) {
	if len(u.UserIDList) > 100 {
		return false, "最多100个"
	}

	return true, ""
}

func (u *UserGetInfo) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *UserGetInfo) SetApiParas(json string) {
}

func (u *UserGetInfo) SetUserIDsList(list []string) {
	u.UserIDList = list
	ids := strings.Join(list, ",")
	u.SetApiOtherPara("userids", ids)
}

func (u *UserGetInfo) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *UserGetInfo) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
