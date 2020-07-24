package wechat

// api url don't need "https" and route, like this "oapi.campus.qq.com"
func NewWeChat(secretId int, secretKey, apiUrl string) *WeChat {
	return &WeChat{
		SecretId:  secretId,
		SecretKey: secretKey,
		ApiUrl:    apiUrl,
	}
}
// 执行签名并发起请求
func (w *WeChat)authAndRequest(orgId, method, action, api string, data RequestData,) (response *Response, err error) {
	url, err := w.Auth(orgId, data, method, api, action)
	if err != nil {
		return nil, err
	}
	return request(url, "GET", data)
}
