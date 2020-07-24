package wechat

type WeChat struct {
	SecretId  int
	SecretKey string
	ApiUrl    string
	Conf *Config
}

// ApiUrl不需要 "https" 和请求的 api 名称, example： "oapi.campus.qq.com"
func NewWeChat(secretId int, secretKey, apiUrl string) *WeChat {
	return &WeChat{
		SecretId:  secretId,
		SecretKey: secretKey,
		ApiUrl:    apiUrl,
	}
}

// 执行签名并发起请求
func (w *WeChat) authAndRequest(orgId, method, action, api string, data RequestData, ) (response *Response, err error) {
	url, err := w.Auth(orgId, data, method, api, action)
	if err != nil {
		return nil, err
	}
	return w.request(url, method, data)
}
