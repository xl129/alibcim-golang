package im

import (
	"encoding/json"
	"fmt"
)

//添加用户
type UserAdd struct {
	UserList []User //`json:"userinfos"`
	Paras    map[string]string
	ResponseData struct {
		Data struct {
			UidSucc struct{ String []string `json:"string"` }  `json:"uid_succ"`
			UidFail struct{ String [] string `json:"string"` } `json:"uid_fail"`
			FailMsg struct{ String [] string `json:"string"` } `json:"fail_msg"`
		} `json:"openim_users_add_response"`
	}
}

func NewUserAdd() *UserAdd {
	return &UserAdd{Paras: make(map[string]string)}
}

func (u *UserAdd) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.users.add"
}

func (u *UserAdd) CheckApiParas() (bool, string) {
	for _, v := range u.UserList {
		if v.UserId == "" {
			return false, "请输入UserId"
		}
		if v.Password == "" {
			return false, "请输入Password"
		}
	}
	if len(u.UserList) > 100 {
		return false, "最多100个"
	}

	return true, ""
}

func (u *UserAdd) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *UserAdd) SetApiParas(json string) {
}

func (u *UserAdd) SetUserList(list []User) {
	bytes, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
		return
	}
	u.UserList = list
	u.SetApiOtherPara("userinfos", string(bytes))
}

func (u *UserAdd) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *UserAdd) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
