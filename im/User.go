package im

type User struct {
	Nick     string `json:"nick"`     //String 可选 king昵称，最大长度64字节
	IconUrl  string `json:"icon_url"` //String 可选 http://xxx.com/xxx头像url，最大长度512字节
	Email    string `json:"email"`    //String 可选 uid@taobao.comemail地址，最大长度128字节
	Mobile   string `json:"mobile"`   //String 可选 18600000000手机号码，最大长度16字节
	TaoBaoId string `json:"taobaoid"` //String 可选 taobaouser淘宝账号，最大长度64字节
	UserId   string `json:"userid"`   //String 必须 imuserim用户名，最大长度64字节
	Password string `json:"password"` //String 必须 xxxxxxim密码，最大长度64字节
	Remark   string `json:"remark"`   //String 可选 demoremark，最大长度128字节
	Extra    string `json:"extra"`    //String 可选 {}扩展字段（Json），最大长度4096字节
	Career   string `json:"career"`   //String 可选 demo职位，最大长度128字节
	Vip      string `json:"vip"`      //String 可选 {}vip（Json），最大长度512字节
	Address  string `json:"address"`  //String 可选 demo地址，最大长度512
	Name     string `json:"name"`     //String 可选 demo名字，最大长度64
	Age      uint8  `json:"age"`      //Number 可选 123年龄
	Gender   string `json:"gender"`   //String 可选 M性别。M: 男。 F：女
	WeChat   string `json:"wechat"`   //String 可选 demo微信，最大长度64字节
	QQ       string `json:"qq"`       //String 可选 demoqq，最大长度20字节
	WeiBo    string `json:"weibo"`    //String 可选 demo微博，最大长度256字节
}
