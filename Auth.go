package wechat

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Params map[string]string
type RequestData map[string]interface{}

func (w *WeChat) Auth(OrgId string, data RequestData, method, apiRouter, action string) (string, error) {
	var params = make(map[string]string)
	params["Action"] = action
	params["Timestamp"] = generateTimestamp()
	params["SecretId"] = strconv.Itoa(w.SecretId)
	params["OrgId"] = OrgId
	params["Nonce"] = generateNonce()
	url := w.ApiUrl + apiRouter
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	sign, err := Sign(w.SecretKey, method, url, params, string(jsonData))
	if err != nil {
		return "", err
	}
	params["Sign"] = sign
	if method == "GET" && data != nil {
		for k, v := range data {
			val, _ := json.Marshal(v)
			params[k] = string(val)
		}
	}

	httpUrl := SpliceUrl(params)
	httpUrl = fmt.Sprintf("https://%s?%s", url, httpUrl)
	return httpUrl, nil
}

func generateNonce() string {
	return strconv.Itoa(rand.Intn(999999))
}

func generateTimestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
