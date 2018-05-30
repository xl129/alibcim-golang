package im

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ResponseData struct {
	Code    int
	Msg     string
	SubCode string
	SubMsg  string
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
	RequestId string `json:"request_id"`
}

type Response struct {
	ErrorResponse `json:"error_response"`
}

func newResponseData() ResponseData {
	return ResponseData{}
}

func (res *ResponseData) addCode(Code int) {
	res.Code = Code
}

func (res *ResponseData) addMsg(Msg string) {
	res.Msg = Msg
}

func (res *ResponseData) addSubCode(SubCode string) {
	res.SubCode = SubCode
}

func (res *ResponseData) addSubMsg(SubMsg string) {
	res.SubMsg = SubMsg
}

//接口请求
func Execute(baseInterface BaseInterface) ResponseData {
	config := NewConfig()
	ok, ApiParas, msg := baseInterface.GetApiParas()
	//fmt.Println(paras)
	if !ok {
		ResData := newResponseData()
		ResData.addCode(5000)
		ResData.addMsg(msg)
		return ResData
	}
	// 根据文档  文件和数组 字节不参与签名
	var signData = make(map[string]string)
	var baseUrl string
	var secretKey string
	var postData = url.Values{}
	query := url.Values{}
	for k, v := range config.Paras {
		if k == "sign" {
			continue
		}
		if k == "url" {
			baseUrl = v
			continue
		}
		if k == "secret_key" {
			secretKey = v
			continue
		}
		query.Add(k, v)
		signData[k] = v
	}
	for k, v := range ApiParas {
		if k == "method" {
			query.Add(k, v)
		}
		if k != "method" {
			postData.Add(k, v)
		}
		signData[k] = v
	}
	sign := Sign(signData, secretKey)
	query.Add("sign", sign)
	queryStr := query.Encode()
	apiUrl := baseUrl + "?" + queryStr

	return httpPostForm(baseInterface, apiUrl, postData)
}

//发起表单提交
func httpPostForm(baseInterface BaseInterface, uri string, values url.Values) ResponseData {
	ResData := newResponseData()
	resp, err := http.PostForm(uri, values)
	if err != nil {
		ResData.addCode(5000)
		ResData.addMsg(err.Error())
		return ResData
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ResData.addCode(5000)
		ResData.addMsg(err.Error())
		return ResData
	}
	response := &Response{ErrorResponse{Code: 0}}
	jErr := json.Unmarshal(body, response)
	// json 解析出错
	if jErr != nil {
		ResData.addCode(5000)
		ResData.addMsg(response.Msg)
		ResData.addSubCode(response.SubCode)
		ResData.addSubMsg(response.SubMsg)
		return ResData
	}
	//表示错误返回解析错误
	if response.Code != 0 {
		ResData.addCode(response.Code)
		ResData.addMsg(response.Msg)
		ResData.addSubCode(response.SubCode)
		ResData.addSubMsg(response.SubMsg)
		return ResData
	}
	ResData.addCode(0)
	//解析返回
	baseInterface.SetResponse(body)
	return ResData
}
