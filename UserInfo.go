package wechat

import "net/http"

type UsersInfos struct {
	Total    int         `json:"Total"`
	DataList []*UserInfo `json:"DataList"`
}
type UserInfo struct {
	OrgUserId string `json:"OrgUserId"`
	OpenUserId string `json:"OpenUserId"`
	RoleId int `json:"RoleId"`
	Name string `json:"Name"`
	Sex int `json:"Sex"`
	DepartmentIds []int `json:"DepartmentIds"`
	Avatar string `json:"Avatar"`
	DepartmentNames []string `json:"DepartmentNames"`
	JoinDate string `json:"JoinDate"`
	TitleIds []int `json:"TitleIds"`
	Status int `json:"Status"`
	CreateTime string `json:"CreateTime"`
	UserNo string `json:"UserNo"`
}
type UsersInfoDetails struct {
	Total    int               `json:"Total"`
	DataList []*UserInfo `json:"DataList"`
}
type UserGroups struct {
	GroupIds []int `json:"GroupIds"`
}
//获取用户的基本信息
func (w *WeChat) GetUsersInfo(orgId string, data RequestData) (*UsersInfos, error) {
	return w.userList(orgId, http.MethodPost,"GetUsersInfo",ORG_USER_API,data)
}
// 根据架构id获取用户信息列表
func (w *WeChat) GetUserListByDepartmentIds(orgId string, data RequestData) (*UsersInfos, error) {
	return w.userList(orgId, http.MethodPost,"GetUserListByDepartmentIds",ORG_USER_API,data)
}

func (w *WeChat) GetUsersInfoDetail(orgId string, data RequestData) (*UsersInfoDetails, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodPost, "GetUsersInfoDetail", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var details UsersInfoDetails
	if err := response.DecodeData(details); err != nil {
		return nil, err
	}
	return &details, nil
}

//该接口比较特殊，即使添加失败了，ErrorCode也是Ok, 请使用FailedIdx进行判断. 这是官方文档的说明
func (w *WeChat) GetUsersInfoUserNo(orgId string, data RequestData) {
	//TODO 官方文档有问题，待官方修改后再完成
}

// 获取教师的其他身份
func (w *WeChat) UserGroups(orgId string, data RequestData) (*UserGroups, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodPost, "UserGroups", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var groups UserGroups
	if err := response.DecodeData(&groups); err != nil {
		return nil, err
	}
	return &groups, err
}

func (w *WeChat)userList(orgId, method, action, api string, data RequestData)(*UsersInfos, error) {
	response, err := w.AuthAndRequest(orgId, method, action, api, data)
	if err != nil {
		return nil, err
	}

	var userInfos UsersInfos
	err = response.DecodeData(&userInfos)

	return &userInfos, err
}
