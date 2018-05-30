package im

import (
	"encoding/json"
	"strconv"
)

// 发送群消息
type TribeSendMsg struct {
	User    ApiUser `json:"user"`
	TribeId uint64  `json:"tribe_id"`
	Msg struct {
		AtFlag     int       `json:"at_flag"`
		AtMembers  []ApiUser `json:"atmembers"`
		CustomPush string    `json:"custom_push"`
		MediaAttrs string    `json:"media_attrs"`
		MsgContent string    `json:"msg_content"`
		MsgType    int       `json:"msg_type"`
		Push       bool      `json:"push"`
	} `json:"msg"`
	Paras map[string]string
	ResponseData struct {
		Data struct {
			Message   string `json:"message"`
			TribeCode int    `json:"tribe_code"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribe_sendmsg_response"`
	}
}

func (u *TribeSendMsg) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribe.sendmsg"
}

func (u *TribeSendMsg) CheckApiParas() (bool, string) {
	if u.User.Uid == "" {
		return false, "缺少User.uid"
	}
	if u.TribeId <= 0 {
		return false, "tribe_id错误"
	}
	return true, ""
}

func (u *TribeSendMsg) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}
	bytes, err := json.Marshal(u.Msg)
	if err != nil {
		return false, data, err.Error()
	}
	u.SetApiMethodName()
	u.SetApiOtherPara("msg", string(bytes))
	return true, u.Paras, ""
}

func (u *TribeSendMsg) SetApiParas(json string) {
}

func (u *TribeSendMsg) SetApiUser(user ApiUser) () {
	u.User = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user", string(bytes))
}

func (u *TribeSendMsg) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeSendMsg) SetMsgAtFlag(AtFlag int) {
	u.Msg.AtFlag = AtFlag
}

func (u *TribeSendMsg) SetMsgAtMembers(members []ApiUser) {
	u.Msg.AtMembers = members
}

func (u *TribeSendMsg) SetMsgCustomPush(d, sound, title string) {
	custom := map[string]string{
		"d":     d,
		"sound": sound,
		"title": title,
	}

	bytes, err := json.Marshal(custom)
	if err != nil {
		return
	}
	u.Msg.CustomPush = string(bytes)
}

// 参考文档
func (u *TribeSendMsg) SetMsgMediaAttrs(json string) {
	u.Msg.MediaAttrs = json
}

func (u *TribeSendMsg) SetMsgMsgContent(str string) {
	u.Msg.MsgContent = str
}

func (u *TribeSendMsg) SetMsgMsgType(int int) {
	u.Msg.MsgType = int
}

func (u *TribeSendMsg) SetMsgPush(bool bool) {
	u.Msg.Push = bool
}

func (u *TribeSendMsg) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeSendMsg) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
