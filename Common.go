package wechat

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	Url "net/url"
	"sort"
	"strings"
	"time"
)

type Config struct {
	AppId     int
	SecretKey string
	ApiAddr   string

	Dial      time.Duration
	Timeout   time.Duration
	KeepAlive time.Duration
	MaxConn   int
	MaxIdle   int

	BackoffInterval time.Duration
	retryCount      int
}

type Response struct {
	ErrorCode ErrCode         `json:"ErrorCode"`
	RequestId string          `json:"RequestId"`
	ErrorMsg  string          `json:"ErrorMsg"`
	Detail    string          `json:"Detail"`
	Data      json.RawMessage `json:"Data"`
}

func (r *Response) DecodeData(v interface{}) error {
	reader := bytes.NewReader(r.Data)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

func Sign(secretKey, method, url string, params Params, body string) (signStr string, err error) {
	paramsStr := SpliceUrl(params)

	rawStr := fmt.Sprintf("%s%s?%s", method, url, paramsStr)
	if (method == http.MethodPost || method == http.MethodPut) && len(body) > 0 {
		rawStr += fmt.Sprintf("&Data=%s", body)
	}

	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(rawStr))

	signStr = Url.QueryEscape(hex.EncodeToString(mac.Sum(nil)))
	return
}

// Splice request params
func SpliceUrl(params Params) string {
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	paramsPairs := make([]string, 0)
	for _, k := range keys {
		paramsPairs = append(paramsPairs, fmt.Sprintf("%s=%s", k, params[k]))
	}
	paramsStr := strings.Join(paramsPairs, "&")

	return paramsStr
}

func checkErrCode(res *Response) bool {
	if res.ErrorCode == ERROR_CODE_OK {
		return true
	}
	return false
}
