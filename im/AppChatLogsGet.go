package im

import (
	"encoding/json"
	"strconv"
)

type ESMessage struct {
	Time   int64   `json:"time"`
	Uuid   int64   `json:"uuid"`
	Type   int     `json:"type"`
	FromId ApiUser `json:"from_id"`
	ToId   ApiUser `json:"to_id"`
	Content struct {
		RoamingMessageItem []RoamingMessageItem `json:"roaming_message_item"`
	} `json:"content"`
}

//应用聊天记录查询
type AppChatLogsGet struct {
	Beg   int64  `json:"beg"`
	End   int64  `json:"end"`
	Count int    `json:"count"`
	Next  string `json:"next"`
	Paras map[string]string
	ResponseData struct {
		Data struct {
			Result struct {
				Messages struct {
					ESMessage [] ESMessage `json:"es_message"`
				} `json:"messages"`
				NextKey string `json:"next_key"`
			} `json:"result"`
			RequestId string `json:"request_id"`
		} `json:"openim_app_chatlogs_get_response"`
	}
}

func (u *AppChatLogsGet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.app.chatlogs.get"
}

func (u *AppChatLogsGet) CheckApiParas() (bool, string) {
	if u.Count <= 0 {
		return false, "缺少Count"
	}
	if u.Beg <= 0 {
		return false, "Begin错误"
	}
	if u.End <= 0 {
		return false, "End错误"
	}
	if u.Beg > u.End {
		return false, "End不能小于Begin"
	}
	return true, " "
}

func (u *AppChatLogsGet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *AppChatLogsGet) SetApiParas(json string) {
}

func (u *AppChatLogsGet) SetCount(count int) () {
	u.Count = count
	u.SetApiOtherPara("count", strconv.Itoa(count))
}

func (u *AppChatLogsGet) SetNext(Next string) () {
	u.Next = Next
	u.SetApiOtherPara("next", Next)
}

func (u *AppChatLogsGet) SetBeg(begin int64) () {
	u.Beg = begin
	u.SetApiOtherPara("beg", strconv.FormatInt(begin, 10))
}

func (u *AppChatLogsGet) SetEnd(end int64) () {
	u.End = end
	u.SetApiOtherPara("end", strconv.FormatInt(end, 10))
}

func (u *AppChatLogsGet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *AppChatLogsGet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
