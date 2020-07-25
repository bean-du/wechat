package wechat

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Title         string  `json:"Title"`
	Content       string  `json:"Content"`
	Users         []*User `json:"Users"`
	Departments   []int   `json:"Departments"`
	AppId         int     `json:"AppId"`
	AppMessageUrl string  `json:"AppMessageUrl"`
}

type User struct {
	OrgUserId   string `json:"OrgUserId"`
	OrgUserName string `json:"OrgUserName"`
	ChildId     string `json:"ChildId"`
	ChildName   string `json:"ChildName"`
}

type MessageList struct {
	PageInfo PageInfo `json:"PageInfo"`
	DataList []*MessageResponse `json:"DataList"`
}
type PageInfo struct {
	Total int `json:"Total"`
	Page int `json:"Page"`
	Size int `json:"Size"`
}
type MessageResponse struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	OrgUserId string `json:"OrgUserId"`
	Name string `json:"Name"`
	Read string `json:"Read"`
	Status int `json:"Status"`
	StatusStr string `json:"StatusStr"`
	SendTime string `json:"SendTime"`
	SendNum int `json:"SendNum"`
	ReadNum int `json:"ReadNum"`
	Timer string `json:"Timer"`
	Tag int `json:"Tag"`
	TagLabel string `json:"TagLabel"`
	Range json.RawMessage `json:"Range"`
}

//发送信息通知
func (w *WeChat)SendMessage(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodPost, "SendMessage", MSG_API, data)
}
//获取消息发送列表
func (w *WeChat)GetMessageList(orgId string, data RequestData) (*MessageList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetMessageList", MSG_API, data)
	if err != nil {
		return nil, err
	}
	res := new(MessageList)
	err = response.DecodeData(res)
	return res, err
}
//获取消息详情
func (w *WeChat)GetMessageDetail(orgId string, data RequestData) (*MessageResponse, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetMessageDetail", MSG_API, data)
	if err != nil {
		return nil, err
	}
	res := new(MessageResponse)
	err = response.DecodeData(res)
	return res, err
}
