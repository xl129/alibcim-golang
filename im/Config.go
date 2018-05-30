package im

import (
	"time"
)

const env = "prod_http"

const prod_http = "http://gw.api.taobao.com/router/rest"

const prod_https = "https://eco.taobao.com/router/rest"

const test_http = "http://gw.api.tbsandbox.com/router/rest"

const test_https = "https://gw.api.tbsandbox.com/router/rest"

const app_key = "24905796"

const secret_key = "2b85dbdcf868ff2918a7e2a23d2aa255"

const sign_method = "hmac"

const version = "2.0"

type Config struct {
	//url            string //请求的地址
	//method         string //	是	API接口名称。
	//app_key        string //	是	TOP分配给应用的AppKey。
	//secret_key     string
	//target_app_key string //	否	被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效。
	//sign_method    string //	是	签名的摘要算法，可选值为：hmac，md5。
	//sign           string //	是	API输入参数签名结果，签名算法介绍请点击这里。
	//session        string //	否	用户登录授权成功后，TOP颁发给应用的授权信息，详细介绍请点击这里。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选。
	//timestamp      string //	是	时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟。
	//format         string //	否	响应格式。默认为xml格式，可选值：xml，json。
	//v              string //	是	API协议版本，可选值：2.0。
	//partner_id     string //	否	合作伙伴身份标识。
	//simplify       bool   //	否	是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false。
	//simplify 不考虑这个值
	Paras map[string]string
}

func (c *Config) setAppKey() {
	c.Paras["app_key"] = app_key
}

func (c *Config) setSecretKey() {
	c.Paras["secret_key"] = secret_key
}

func (c *Config) setTimestamp() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	c.Paras["timestamp"] = now.In(local).Format(format)
}

func (c *Config) setSignMethod() {
	c.Paras["sign_method"] = sign_method
}

func (c *Config) setVersion() {
	c.Paras["v"] = version
}

func (c *Config) setUrl() {
	url := ""
	switch env {
	case "test_http":
		url = test_http
		break
	case "test_https":
		url = test_https
		break
	case "prod_http":
		url = prod_http
		break
	case "prod_https":
		url = prod_https
		break
	}
	c.Paras["url"] = url
}

func (c *Config) setFormat() {
	c.Paras["format"] = "json"
}

var (
	ConfigSingle *Config
)

//new 一个 配置文件
func NewConfig() *Config {
	if ConfigSingle == nil {
		ConfigSingle = &Config{Paras: make(map[string]string)}
		ConfigSingle.setUrl()
		ConfigSingle.setAppKey()
		ConfigSingle.setSecretKey()
		ConfigSingle.setSignMethod()
		ConfigSingle.setVersion()
		ConfigSingle.setFormat()
	}
	ConfigSingle.setTimestamp()

	return ConfigSingle
}
