package im

import (
	"encoding/json"
)

//标准消息发送
type IMMsgPush struct {
	IMMsg
	Paras map[string]string
	ResponseData struct {
		Data struct {
			MsgId     uint64 `json:"msgid"`
			RequestId string `json:"request_id"`
		} `json:"openim_immsg_push_response"`
	}
}

func NewIMMsgPush() *IMMsgPush {
	return &IMMsgPush{Paras: make(map[string]string)}
}

func (u *IMMsgPush) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.immsg.push"
}

func (u *IMMsgPush) CheckApiParas() (bool, string) {
	if u.FromUser == "" {
		return false, "缺少FromUser"
	}
	if len(u.ToUsers) == 0 {
		return false, "ToUsers错误"
	}
	if u.Context == "" {
		return false, "Context错误"
	}
	return true, ""
}

func (u *IMMsgPush) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}
	if u.ToAppKey == "" {
		u.SetToAppKey(app_key)
	}
	bytes, err := json.Marshal(u.IMMsg)
	if err != nil {
		return false, data, err.Error()
	}
	u.SetApiMethodName()
	u.SetApiOtherPara("immsg", string(bytes))
	return true, u.Paras, ""
}

func (u *IMMsgPush) SetApiParas(json string) {
}

func (u *IMMsgPush) SetFromUser(uid string) () {
	u.FromUser = uid
}

func (u *IMMsgPush) SetToAppKey(app_key string) {
	u.ToAppKey = app_key
}

func (u *IMMsgPush) SetToUsers(uids []string) {
	u.ToUsers = uids
}

func (u *IMMsgPush) SetMsgType(MsgType int) {
	u.MsgType = MsgType
}

func (u *IMMsgPush) SetContext(Context string) {
	u.Context = Context
}

func (u *IMMsgPush) SetFromTaoBao(FromTaoBao int) {
	u.FromTaoBao = FromTaoBao
}

func (u *IMMsgPush) SetMediaAttr(string string) {
	u.MediaAttr = string
}

func (u *IMMsgPush) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *IMMsgPush) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
