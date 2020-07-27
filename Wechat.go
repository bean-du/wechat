package wechat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type WeChat struct {
	SecretId  int
	SecretKey string
	ApiUrl    string
	Conf      *Config
}

// ApiUrl不需要 "https" 和请求的 api 名称, example： "oapi.campus.qq.com"
// config 为网络请求已经重试配置
func NewWeChat(conf *Config) *WeChat {
	if conf.retryCount == 0 {
		conf.retryCount = 2
	}
	return &WeChat{
		SecretId:  conf.AppId,
		SecretKey: conf.SecretKey,
		ApiUrl:    conf.ApiAddr,
		Conf:      conf,
	}
}

func (w *WeChat) SetRetry(retry int) {
	w.Conf.retryCount = retry
}

// 执行签名并发起请求
func (w *WeChat) AuthAndRequest(orgId, method, action, api string, data RequestData, ) (response *Response, err error) {
	url, err := w.Auth(orgId, data, method, api, action)
	if err != nil {
		return nil, err
	}
	return w.request(url, method, data)
}

func (w *WeChat) request(url, method string, data RequestData) (response *Response, err error) {
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
	if err != nil {
		return nil, err
	}

	if !checkErrCode(response) {
		return nil, errors.New(response.ErrorMsg + response.Detail)
	}
	return
}

