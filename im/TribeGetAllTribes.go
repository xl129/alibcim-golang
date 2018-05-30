package im

import (
	"encoding/json"
	"strings"
	"fmt"
)

//获取用户群列表
type TribeGetAllTribes struct {
	User       ApiUser `json:"user"`
	TribeTypes [] string
	Paras      map[string]string
	ResponseData struct {
		Data struct {
			TribeInfoList struct{
				TribeInfo []Tribe `json:"tribe_info"`
			} `json:"tribe_info_list"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_getalltribes_response"`
	}
}

func (u *TribeGetAllTribes) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.getalltribes"
}

func (u *TribeGetAllTribes) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if len(u.TribeTypes) == 0 {
		return false, "TribeTypes错误"
	}
	return true, ""
}

func (u *TribeGetAllTribes) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeGetAllTribes) SetApiParas(json string) {
}

func (u *TribeGetAllTribes) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeGetAllTribes) SetTribeTypes(tribe_types []string) () {
	u.TribeTypes = tribe_types
	u.SetApiOtherPara("tribe_types", strings.Join(tribe_types, ","))
}

func (u *TribeGetAllTribes) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeGetAllTribes) SetResponse(body []byte) {
	fmt.Println(string(body))
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
