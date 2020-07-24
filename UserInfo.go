package wechat

var userApi = "/v2/user"

type UsersInfos struct {
	Total int `json:"Total"`
	DataList []*UserInfo `json:"DataList"`
}

type UserInfo struct {

}

type UsersInfoDetails struct {
	Total int `json:"Total"`
	DataList []*UserInfoDetail `json:"DataList"`
}

type UserInfoDetail struct {

}

func (w *WeChat)GetUsersInfo(orgId string, data RequestData) (*UsersInfos, error) {
	url, err := w.Auth(orgId, data, "POST", userApi, "GetOrgInfo")
	if err != nil {
		return nil, err
	}
	response, err := Request(url,"GET", data)
	if err != nil {
		return nil, err
	}

	var userInfos UsersInfos
	if err := response.DecodeData(&userInfos); err != nil {
		return nil, err
	}
	return &userInfos, nil
}

func (w *WeChat)GetUsersInfoDetail(orgId string, data RequestData) (*UsersInfoDetails, error) {
	url, err := w.Auth(orgId, data, "POST", userApi, "GetUsersInfoDetail")
	if err != nil {
		return nil, err
	}

	response, err := Request(url,"GET", data)
	if err != nil {
		return nil, err
	}

	var details UsersInfoDetails
	if err := response.DecodeData(details); err != nil {
		return  nil, err
	}
	return &details, nil
}



