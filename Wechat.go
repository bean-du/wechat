package wechat

// api url don't need "https" and route, like this "oapi.campus.qq.com"
func NewWeChat(secretId int, secretKey, apiUrl string) *WeChat {
	return &WeChat{
		SecretId:  secretId,
		SecretKey: secretKey,
		ApiUrl:    apiUrl,
	}
}
