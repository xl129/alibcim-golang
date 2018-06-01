package im

import (
	"encoding/json"
	"time"
)

//账号的聊天关系
type RelationsGet struct {
	User    ApiUser `json:"user"`
	BegDate string  `json:"beg_date"`
	EndDate string  `json:"end_date"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			Users struct {
				OpenImUser []User `json:"open_im_user"`
			} `json:"users"`
			RequestId string `json:"request_id"`
		} `json:"openim_relations_get_response"`
	}
}

func NewRelationsGet() *RelationsGet {
	return &RelationsGet{Paras: make(map[string]string)}
}

func (u *RelationsGet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.relations.get"
}

func (u *RelationsGet) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.BegDate == "" {
		return false, "BegDate错误"
	}
	if u.EndDate == "" {
		return false, "EndDate错误"
	}
	// 获取一个月前的时间戳
	now := time.Now()
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return false, err.Error()
	}
	prv_month := now.In(local).AddDate(0, -1, 0).Unix()

	format := "20060102"
	beg_time, err := time.ParseInLocation(format, u.BegDate, local)
	if err != nil {
		return false, err.Error()
	}
	unix_beg_time := beg_time.Unix()
	if unix_beg_time < prv_month {
		return false, "BegDate超过一个月"
	}
	end_time, err := time.ParseInLocation(format, u.EndDate, local)
	if err != nil {
		return false, err.Error()
	}
	unix_end_time := end_time.Unix()
	if unix_end_time < unix_beg_time {
		return false, "EndDate不能小于BegDate"
	}
	return true, " "
}

func (u *RelationsGet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *RelationsGet) SetApiParas(json string) {
}

func (u *RelationsGet) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *RelationsGet) SetBegDate(beg_date string) () {
	u.BegDate = beg_date
	u.SetApiOtherPara("beg_date", beg_date)
}

func (u *RelationsGet) SetEndDate(end_date string) () {
	u.EndDate = end_date
	u.SetApiOtherPara("end_date", end_date)
}

func (u *RelationsGet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *RelationsGet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
