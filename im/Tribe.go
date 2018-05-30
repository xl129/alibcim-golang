package im

type Tribe struct {
	TribeId   uint64 `json:"tribe_id"`   //123群ID
	Icon      string `json:"icon"`       //demo群头像URL地址
	CheckMode int    `json:"check_mode"` //123群验证模式
	TribeType int    `json:"tribe_type"` //123群类型 群类型有两种tribe_type = 0 普通群 普通群有管理员角色，对成员加入有权限控制tribe_type = 1 讨论组 讨论组没有管理员，不能解散
	Name      string `json:"name"`       //demo群名称
	RecvFlag  int    `json:"recv_flag"`  //123群接收标记
	Notice    string `json:"notice"`     //demo群公告
}
