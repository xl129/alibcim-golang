package im

import (
	"encoding/json"
)

//设置开发环境证书
type IOSCertSandBoxSet struct {
	Cert     string `json:"cert"`
	Password string `json:"password"`
	Paras    map[string]string
	ResponseData struct {
		Data struct {
			Code      int    `json:"code"`
			RequestId string `json:"request_id"`
		} `json:"openim_ioscert_sandbox_set_response"`
	}
}

func (u *IOSCertSandBoxSet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.ioscert.sandbox.set"
}

func (u *IOSCertSandBoxSet) CheckApiParas() (bool, string) {
	if u.Cert == "" {
		return false, "缺少Cert"
	}
	return true, ""
}

func (u *IOSCertSandBoxSet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *IOSCertSandBoxSet) SetApiParas(json string) {
}

func (u *IOSCertSandBoxSet) SetCert(cert string) () {
	u.Cert = cert
	u.SetApiOtherPara("notice", cert)
}

func (u *IOSCertSandBoxSet) SetPassword(password string) () {
	u.Password = password
	u.SetApiOtherPara("password", password)
}

func (u *IOSCertSandBoxSet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *IOSCertSandBoxSet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
