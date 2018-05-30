package im

import (
	"encoding/json"
	"strconv"
)

//聊天记录查询接口
type ChatLogsGet struct {
	User1   ApiUser `json:"user1"`
	User2   ApiUser `json:"user2"`
	Begin   int64   `json:"begin"`
	End     int64   `json:"end"`
	Count   int     `json:"count"`
	NextKey string  `json:"next_key"`
	Paras   map[string]string
	ResponseData struct {
		Data struct {
			Result struct {
				Messages struct {
					RoamingMessage [] RoamingMessage `json:"roaming_message"`
				} `json:"messages"`
				NextKey string `json:"next_key"`
			} `json:"result"`
			RequestId string `json:"request_id"`
		} `json:"openim_chatlogs_get_response"`
	}
}

func (u *ChatLogsGet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.chatlogs.get"
}

func (u *ChatLogsGet) CheckApiParas() (bool, string) {
	if u.User1.Uid == "" {
		return false, "缺少User1.uid"
	}
	if u.User2.Uid == "" {
		return false, "缺少User2.uid"
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

func (u *ChatLogsGet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *ChatLogsGet) SetApiParas(json string) {
}

func (u *ChatLogsGet) SetApiUser1(user ApiUser) () {
	u.User1 = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user1", string(bytes))
}

func (u *ChatLogsGet) SetApiUser2(user ApiUser) () {
	u.User2 = user
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	u.SetApiOtherPara("user2", string(bytes))
}

func (u *ChatLogsGet) SetCount(count int) () {
	u.Count = count
	u.SetApiOtherPara("count", strconv.Itoa(count))
}

func (u *ChatLogsGet) SetNextKey(NextKey string) () {
	u.NextKey = NextKey
	u.SetApiOtherPara("next_key", NextKey)
}

func (u *ChatLogsGet) SetBegin(begin int64) () {
	u.Begin = begin
	u.SetApiOtherPara("begin", strconv.FormatInt(begin, 10))
}

func (u *ChatLogsGet) SetEnd(end int64) () {
	u.End = end
	u.SetApiOtherPara("end", strconv.FormatInt(end, 10))
}

func (u *ChatLogsGet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *ChatLogsGet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
