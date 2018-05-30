package im

import (
	"encoding/json"
	"strconv"
)

type TribeMessage struct {
	Content struct {
		MessageItem []RoamingMessageItem `json:"message_item"`
	} `json:"content"`
	FromId ApiUser `json:"from_id"`
	Time   int64   `json:"time"`
	Uuid   int64   `json:"uuid"`
	Type   int     `json:"type"`
}

//群聊天记录导出接口
type TribeLogsGet struct {
	TribeId uint64 `json:"tribe_id"`
	Begin   int64  `json:"begin"`
	End     int64  `json:"end"`
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			RetCode int    `json:"retCode"`
			Message string `json:"message"`
			Reason  string `json:"reason"`
			Succ    string `json:"succ"`
			Data struct {
				Messages struct {
					TribeMessage []TribeMessage `json:"tribe_message"`
				} `json:"messages"`
				NextKey string `json:"next_key"`
			} `json:"data"`
			RequestId string `json:"request_id"`
		} `json:"openim_tribelogs_get_response"`
	}
}

func (u *TribeLogsGet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.tribelogs.get"
}

func (u *TribeLogsGet) CheckApiParas() (bool, string) {
	if u.TribeId <= 0 {
		return false, "缺少TribeId"
	}
	if u.Count <= 0 {
		return false, "缺少Count"
	}
	if u.Begin <= 0 {
		return false, "Begin错误"
	}
	if u.End <= 0 {
		return false, "End错误"
	}
	if u.Begin > u.End {
		return false, "End不能小于Begin"
	}
	return true, " "
}

func (u *TribeLogsGet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *TribeLogsGet) SetApiParas(json string) {
}

func (u *TribeLogsGet) SetTribeId(tribe_id uint64) () {
	u.TribeId = tribe_id
	u.SetApiOtherPara("tribe_id", strconv.FormatUint(tribe_id, 10))
}

func (u *TribeLogsGet) SetCount(count int) () {
	u.Count = count
	u.SetApiOtherPara("count", strconv.Itoa(count))
}

func (u *TribeLogsGet) SetNext(Next string) () {
	u.Next = Next
	u.SetApiOtherPara("next", Next)
}

func (u *TribeLogsGet) SetBegin(begin int64) () {
	u.Begin = begin
	u.SetApiOtherPara("begin", strconv.FormatInt(begin, 10))
}

func (u *TribeLogsGet) SetEnd(end int64) () {
	u.End = end
	u.SetApiOtherPara("end", strconv.FormatInt(end, 10))
}

func (u *TribeLogsGet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *TribeLogsGet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
