package wechat

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	Url "net/url"
	"sort"
	"strings"
)

type WeChat struct {
	SecretId  int
	SecretKey string
	ApiUrl    string
}

type Response struct {
	ErrorCode string          `json:"ErrorCode"`
	RequestId string          `json:"RequestId"`
	ErrorMsg  string          `json:"ErrorMsg"`
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

func RequestApi(url string, method string, data *[]byte) ([]byte, error) {
	var body io.Reader

	if data != nil {
		body = bytes.NewReader(*data)
	}
	client := http.Client{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println("make new request error : ", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		log.Println("do request error: ", err)
		return nil, err
	}
	defer response.Body.Close()

	var res []byte
	res, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Sign(secret_key, method, url string, params Params, body string) (signStr string, err error) {
	paramsStr :=SpliceUrl(params)

	rawStr := fmt.Sprintf("%s%s?%s", method, url, paramsStr)
	if (method == "POST" || method == "PUT") && len(body) > 0 {
		rawStr += fmt.Sprintf("&Data=%s", body)
	}
	mac := hmac.New(sha1.New, []byte(secret_key))
	mac.Write([]byte(rawStr))
	hash := mac.Sum(nil)
	b16encoded := hex.EncodeToString(hash)
	signStr = Url.QueryEscape(b16encoded)
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

func request(url, method string, data RequestData, ) (response *Response, err error)  {
	var  body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	res, err := RequestApi(url, method, &body)
	if err != nil {
		return nil, err
	}

	response = new(Response)
	if err := decode(res, response); err != nil {
		return nil, err
	}
	if !checkErrCode(response) {
		return nil, errors.New(response.ErrorMsg)
	}
	return
}

func decode(data []byte, v interface{})  error {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

func checkErrCode(res *Response) bool {
	if res.ErrorCode == "OK" {
		return true
	}
	return false
}
