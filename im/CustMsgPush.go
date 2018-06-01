package im

import (
	"encoding/json"
	"unicode/utf8"
)

//推送自定义openim消息
type CustMsgPush struct {
	CustMsg
	Paras map[string]string
	ResponseData struct {
		Data struct {
			MsgId     uint64 `json:"msgid"`
			RequestId string `json:"request_id"`
		} `json:"openim_custmsg_push_response"`
	}
}

func NewCustMsgPush() *CustMsgPush {
	return &CustMsgPush{Paras: make(map[string]string)}
}

func (u *CustMsgPush) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.custmsg.push"
}

func (u *CustMsgPush) CheckApiParas() (bool, string) {
	if u.FromUser == "" {
		return false, "缺少FromUser"
	}
	if len(u.ToUsers) == 0 {
		return false, "ToUsers错误"
	}
	if u.Summary == "" {
		return false, "Summary错误"
	}
	if u.Data == "" {
		return false, "Data错误"
	}
	if utf8.RuneCountInString(u.Aps)+utf8.RuneCountInString(u.ApnsParam) > 200 {
		return false, "Aps+ApnsParam > 200 错误"
	}
	return true, ""
}

func (u *CustMsgPush) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}
	if u.ToAppKey == "" {
		u.SetToAppKey(app_key)
	}
	bytes, err := json.Marshal(u.CustMsg)
	if err != nil {
		return false, data, err.Error()
	}
	u.SetApiMethodName()
	u.SetApiOtherPara("custmsg", string(bytes))
	return true, u.Paras, ""
}

func (u *CustMsgPush) SetApiParas(json string) {
}

func (u *CustMsgPush) SetFromUser(uid string) () {
	u.FromUser = uid
}

func (u *CustMsgPush) SetToAppKey(app_key string) {
	u.ToAppKey = app_key
}

func (u *CustMsgPush) SetToUsers(uids []string) {
	u.ToUsers = uids
}

func (u *CustMsgPush) SetSummary(summary string) {
	u.Summary = summary
}

// 自定义消息格式
func (u *CustMsgPush) SetData(data string) {
	u.Data = data
}

func (u *CustMsgPush) SetAps(aps string) {
	var alert = map[string]string{
		"alert": aps,
	}
	bytes, err := json.Marshal(alert)
	if err != nil {
		return
	}
	u.Aps = string(bytes)
}

func (u *CustMsgPush) SetApnsParam(param string) {
	u.ApnsParam = param
}

func (u *CustMsgPush) SetInvisible(invisible int) {
	u.Invisible = invisible
}

func (u *CustMsgPush) SetFromNick(from_nick string) {
	u.FromNick = from_nick
}

func (u *CustMsgPush) SetFromTaoBao(from_tao_bao int) {
	u.FromTaoBao = from_tao_bao
}

func (u *CustMsgPush) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *CustMsgPush) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
