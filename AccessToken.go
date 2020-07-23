package wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Data struct {
	AccessToken string  `json:"AccessToken"`
	ExpireIn    int     `json:"ExpireIn"`
	Session     Session `json:"Session"`
}

type Session struct {
	UserName   string          `json:"UserName"`
	OpenUserId string          `json:"OpenUserId"`
	OrgUserId  string          `json:"OrgUserId"`
	RoleId     int             `json:"RoleId"`
	OrgId      int             `json:"OrgId"`
	ExtData    json.RawMessage `json:"ExtData"`
}

func (w *WeChat) GetAccessToken(userCode string) (*Response, error) {
	apiUrl := fmt.Sprintf("%s?Action=GetAccessTokenByCode&SecretId=%d&SecretKey=%s&UserCode=%s", w.ApiUrl, w.SecretId, w.SecretKey, userCode)
	response, err := RequestApi(apiUrl, "GET", nil)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(response)
	decoder := json.NewDecoder(reader)
	var res Response
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (w *WeChat) FlashToken(c context.Context, userCode, currentToken, string, expire time.Duration, result chan string) {
	go func() {
		timer := time.NewTimer((expire - 1) * time.Second)
		for {
			select {
			case <-timer.C:
				url := fmt.Sprintf("%sRefreshAccessToken&SecretId=%d&SecretKey=%s&CurrentAccessToken=%s",w.ApiUrl, w.SecretId, w.SecretKey, currentToken)
				// flash token
				res, err := RequestApi(url, "GET", nil)
				if err != nil {
					log.Println("get AccessToken Api  failed, error info: ", err.Error())
					goto EXIT
				}
				reader := bytes.NewReader(res)
				decoder := json.NewDecoder(reader)
				var response *Response
				if err := decoder.Decode(response); err != nil {
					log.Println("get AccessToken Api  failed, error info: ", err.Error())
					goto EXIT
				}
				var data *Data
				if err := response.DecodeData(data); err != nil {
					log.Println("get AccessToken Api  failed, error info: ", err.Error())
					goto EXIT
				}
				result <- data.AccessToken
			case <-c.Done():
				goto EXIT
			default:
				time.Sleep((expire - 1) * time.Second)
			}
		}
	EXIT:
	}()
}
