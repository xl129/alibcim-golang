package im

type CustMsg struct {
	FromUser   string   `json:"from_user"`   // 必须 user_sender发送方userid
	ToAppKey   string   `json:"to_appkey"`   // 可选 0接收方appkey，不填默认是发送方appkey，如需跨app发送，需要走审批流程，请联系技术答疑
	ToUsers    []string `json:"to_users"`    // [] 必须 ["user1","user2"]接受者userid列表，单次发送用户数小于100
	Summary    string   `json:"summary"`     // 必须 客户端最近消息里面显示的消息摘要客户端最近消息里面显示的消息摘要
	Data       string   `json:"data"`        // 必须 push payload发送的自定义数据，sdk默认无法解析消息，该数据需要客户端自己解析
	Aps        string   `json:"aps"`         // 可选 {"alert":"ios apns push"}apns推送时，里面的aps结构体json字符串，aps.alert为必填字段。本字段为可选，若为空，则表示不进行apns推送。aps.size() + apns_param.size() < 200
	ApnsParam  string   `json:"apns_param"`  // 可选 apns推送的附带数据apns推送的附带数据。客户端收到apns消息后，可以从apns结构体的"d"字段中取出该内容。aps.size() + apns_param.size() < 200
	Invisible  int      `json:"invisible"`   // 可选 0 默认值：0 表示消息是否在最近会话列表里面展示。如果为1，则消息不在列表展示，可以认为服务端透明的给客户端下法了一个数据，ios的提示任然通过aps字段控制
	FromNick   string   `json:"from_nick"`   // 可选 sender_nick可以指定发送方的显示昵称，默认为空，自动使用发送方用户id作为nick
	FromTaoBao int      `json:"from_taobao"` // 可选 0 默认值：0 如果为1，则表示发送方是一个淘宝账号，该账号必须是本appkey绑定过的客服账号，并且只能给本appkey的用户发消息。通过该参数可以从服务端发起一个客服到用户的会话。
}
