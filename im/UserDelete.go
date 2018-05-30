package im

import (
	"encoding/json"
	"strings"
)

//删除用户
type UserDelete struct {
	UserIDList []string
	Paras      map[string]string
	ResponseData struct {
		Data struct {
			Result struct{ String []string `json:"string"` } `json:"result"` //多个好像也是返回的一个ok
		} `json:"openim_users_delete_response"`
	}
}

func (u *UserDelete) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.users.delete"
}

func (u *UserDelete) CheckApiParas() (bool, string) {
	if len(u.UserIDList) > 100 {
		return false, "最多100个"
	}

	return true, ""
}

func (u *UserDelete) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *UserDelete) SetApiParas(json string) {
}

func (u *UserDelete) SetUserIDsList(list []string) {
	u.UserIDList = list
	ids := strings.Join(list, ",")
	u.SetApiOtherPara("userids", ids)
}

func (u *UserDelete) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *UserDelete) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
