package wechat


type OrgInfo struct {
	Name         string  `json:"Name"`
	Logo         string  `json:"Logo"`
	Type         OrgType `json:"Type"`
	Code         string  `json:"Code"`
	Country      string  `json:"Country"`
	Province     string  `json:"Province"`
	City         string  `json:"City"`
	Area         string  `json:"Area"`
	ProvinceCode string  `json:"ProvinceCode"`
	CityCode     string  `json:"CityCode"`
	AreaCode     string  `json:"AreaCode"`
	Level        int     `json:"Level"`
}

type OrgAdmins struct {
	Total    int         `json:"Total"`
	DataList []*OrgAdmin `json:"DataList"`
}

type OrgAdmin struct {
	OrgUserId     string   `json:"OrgUserId"`
	OpenUserId    string   `json:"OpenUserId"`
	Name          string   `json:"Name"`
	RoleId        int      `json:"RoleId"`
	DepartmentIds []int    `json:"DepartmentIds"`
	AdminType     []string `json:"AdminType"`
}

type OrgTitles struct {
	Total    int         `json:"Total"`
	DataList []*OrgTitle `json:"DataList"`
}

type OrgTitle struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type OrgRelationsList struct {
	DataList []*OrgRelation `json:"data_list"`
}

type OrgRelation struct {
	OrgId int    `json:"OrgId"`
	Name  string `json:"Name"`
	Type  int    `json:"Type"`
	Logo  string `json:"Logo"`
}

func (w *WeChat) GetOrgInfo(orgId string, data RequestData) (*OrgInfo, error) {
	response, err := w.authAndRequest(orgId, "GET", "GetOrgInfo", ORG_USER_API, nil)
	if err != nil {
		return nil, err
	}

	var orgInfo OrgInfo
	if err := response.DecodeData(&orgInfo); err != nil {
		return nil, err
	}
	return &orgInfo, nil
}

func (w *WeChat) GetOrgAdmins(orgId string) (*OrgAdmins, error) {
	response, err := w.authAndRequest(orgId, "GET", "GetOrgAdmins", ORG_USER_API, nil)
	if err != nil {
		return nil, err
	}

	var orgAdmins *OrgAdmins
	if err := response.DecodeData(&orgAdmins); err != nil {
		return nil, err
	}
	return orgAdmins, nil
}

func (w *WeChat) GetOrgTitles(orgId string) (*OrgTitles, error) {
	response, err := w.authAndRequest(orgId, "GET", "GetOrgTitles", ORG_USER_API, nil)
	if err != nil {
		return nil, err
	}

	var orgTitles OrgTitles
	if err := response.DecodeData(&orgTitles); err != nil {
		return nil, err
	}
	return &orgTitles, nil
}

// orgType 为请求关系的类型 2：上级单位 4：学校
func (w *WeChat) GetOfficeRelationsList(orgType int, orgId string) (*OrgRelationsList, error) {
	data := RequestData{"Type": orgType}
	response, err := w.authAndRequest(orgId, "GET", "GetOrgTitles", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var orgRelationList OrgRelationsList
	if err := response.DecodeData(&orgRelationList); err != nil {
		return nil, err
	}
	return &orgRelationList, nil
}
