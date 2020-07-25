package wechat

import "net/http"

type DepartmentList struct {
	Total       int           `json:"Total"`
	Departments []*Department `json:"Departments"`
}

// 具体字段说明请到官网查看
// https://developer.campus.qq.com/docs/%E5%BC%80%E6%94%BEAPI%E6%96%87%E6%A1%A3/%E7%AD%BE%E5%90%8D%E9%89%B4%E6%9D%83%E7%B1%BB%E6%8E%A5%E5%8F%A3/%E7%BB%84%E7%BB%87%E6%9E%B6%E6%9E%84%E4%BF%A1%E6%81%AF%E6%9F%A5%E8%AF%A2/%E8%8E%B7%E5%8F%96%E7%BB%84%E7%BB%87%E6%9E%B6%E6%9E%84%E5%88%97%E8%A1%A8.html
type Department struct {
	DepartmentId   int    `json:"DepartmentId"`
	Name           string `json:"Name"`
	Level          int    `json:"Level"`
	ParentId       int    `json:"ParentId"`
	FullPath       string `json:"FullPath"`
	UsersTotal     int    `json:"UsersTotal"`
	Tag            int    `json:"Tag"`
	StandardGrade  int    `json:"StandardGrade"`
	Extra          string `json:"Extra"`
	DepartmentType int    `json:"DepartmentType"`
	Code           string `json:"Code"`
}

//获取指定类型的全量组织架构列表
func (w *WeChat) GetDepartmentList(orgId string, data RequestData) (list *DepartmentList, err error) {
	return w.departmentOpt(orgId, http.MethodGet, "GetDepartmentList", ORG_USER_API, data)
}

//获取组织架构年级、班级等列表
func (w *WeChat) GetDepartmentListByTag(orgId string, data RequestData) (list *DepartmentList, err error) {
	return w.departmentOpt(orgId, http.MethodGet, "GetDepartmentListByTag", ORG_USER_API, data)
}

//获取指定组织架构详细信息列表
func (w *WeChat) GetBatchDepartment(orgId string, data RequestData) (list *DepartmentList, err error) {
	return w.departmentOpt(orgId, http.MethodGet, "GetBatchDepartment", ORG_USER_API, data)
}

//通过组织架构id，获取其成员列表
func (w *WeChat) GetDptUsers(orgId string, data RequestData) (*UsersInfos, error) {
	return w.userList(orgId, http.MethodGet, "GetDptUsers", ORG_USER_API, data)
}

//获取指定组织架构详细信息
func (w *WeChat) GetDepartment(orgId string, data RequestData) (*Department, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetDepartment", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}
	d := new(Department)
	err = response.DecodeData(d)
	return d, err
}

func (w *WeChat) departmentOpt(orgId, method, action, api string, data RequestData) (list *DepartmentList, err error) {
	response, err := w.AuthAndRequest(orgId, method, action, api, data)
	if err != nil {
		return
	}
	list = new(DepartmentList)
	err = response.DecodeData(list)
	return
}

//添加组织架构，支持批量，由根节点递归创建。创建单个节点时只需传根节点即可
func (w *WeChat) AddDepartments(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodPost, "AddDepartments", ORG_USER_API, data)
}

// 根据id更新组织架构，支持批量
func (w *WeChat) UpdateDepartments(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodPost, "UpdateDepartments", ORG_USER_API, data)
}

//删除指定组织架构
func (w *WeChat) DeleteDepartment(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodPost, "DeleteDepartment", ORG_USER_API, data)
}
