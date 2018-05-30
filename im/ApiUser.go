package im

type ApiUser struct {
	AppKey        string `json:"app_key"`
	TaoBaoAccount bool   `json:"taobao_account"`
	Uid           string `json:"uid"`
}

type TribeUser struct {
	ApiUser
	Role string `json:"role"`
}

// 外部调用这个方法创建接口请求对象 app_key要用应用的app_key 不然报错
func NewApiUser(uid string) ApiUser {
	return ApiUser{AppKey: app_key, Uid: uid}
}
