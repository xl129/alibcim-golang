package im

import (
	"encoding/json"
)

//设置ios证书
type IOSCertProductionSet struct {
	Cert     string `json:"cert"`
	Password string `json:"password"`
	Paras    map[string]string
	ResponseData struct {
		Data struct {
			Code      int    `json:"code"`
			RequestId string `json:"request_id"`
		} `json:"openim_ioscert_production_set_response"`
	}
}

func NewIOSCertProductionSet() *IOSCertProductionSet  {
	return &IOSCertProductionSet{Paras: make(map[string]string)}
}

func (u *IOSCertProductionSet) SetApiMethodName() {
	u.Paras["method"] = "taobao.openim.ioscert.production.set"
}

func (u *IOSCertProductionSet) CheckApiParas() (bool, string) {
	if u.Cert == "" {
		return false, "缺少Cert"
	}
	if u.Password == "" {
		return false, "缺少Password"
	}
	return true, ""
}

func (u *IOSCertProductionSet) GetApiParas() (bool, map[string]string, string) {
	var data = make(map[string]string)
	ok, msg := u.CheckApiParas()
	if !ok {
		return false, data, msg
	}

	u.SetApiMethodName()
	return true, u.Paras, ""
}

func (u *IOSCertProductionSet) SetApiParas(json string) {
}

func (u *IOSCertProductionSet) SetCert(cert string) () {
	u.Cert = cert
	u.SetApiOtherPara("notice", cert)
}

func (u *IOSCertProductionSet) SetPassword(password string) () {
	u.Password = password
	u.SetApiOtherPara("password", password)
}

func (u *IOSCertProductionSet) SetApiOtherPara(key string, value string) {
	u.Paras[key] = value
}

func (u *IOSCertProductionSet) SetResponse(body []byte) {
	err := json.Unmarshal(body, &u.ResponseData)
	if err == nil {
		return
	}
}
