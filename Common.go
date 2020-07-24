package wechat

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	Url "net/url"
	"sort"
	"strings"
	"time"
)

type Config struct {
	Dial            time.Duration
	Timeout         time.Duration
	KeepAlive       time.Duration
	MaxConn        int
	MaxIdle         int
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

func (w *WeChat)request(url, method string, data RequestData) (response *Response, err error) {
	response = new(Response)
	client := NewHttpClient(w.Conf)

	ctx, cancel := context.WithTimeout(context.Background(), w.Conf.Timeout)
	defer cancel()

	switch method {
	case http.MethodPost:
		err = client.Post(ctx, url, MIMEJSON, nil, data, response)
	case http.MethodGet:
		err = client.Get(ctx, url, nil, response)
	}
	if  err != nil {
		return nil, err
	}

	if !checkErrCode(response) {
		return nil, errors.New(response.ErrorMsg + response.Detail)
	}
	return
}

func checkErrCode(res *Response) bool {
	if res.ErrorCode == ERROR_CODE_OK {
		return true
	}
	return false
}
