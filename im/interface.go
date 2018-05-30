package im

type BaseInterface interface {
	// 设置当前请求的方法
	SetApiMethodName()
	// 参数检查
	CheckApiParas() (bool, string)
	//获取参数
	GetApiParas() (bool, map[string]string, string)
	//设置参数
	SetApiParas(value string)
	//设置额外参数
	SetApiOtherPara(key string, value string)
	//解析返回
	SetResponse([]byte)
}
