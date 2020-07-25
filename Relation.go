package wechat

import (
	"net/http"
)

type RelationDataList struct {
	DataList []*Relation `json:"DataList"`
}
type Relation struct {
	OpenUserId string `json:"OpenUserId"`
	OrgUserId string `json:"OrgUserId"`
	RoleId int `json:"RoleId"`
	Name string `json:"Name"`
	Relation int `json:"Relation"`
}

type TeacherClassList struct {
	Total int `json:"Total"`
	DataList []*ClassTeacher `json:"DataList"`
}

type ClassTeacher struct {
	OpenUserId string `json:"OpenUserId"`
	OrgUserId string `json:"OrgUserId"`
	DepartmentId int `json:"DepartmentId"`
	ClassName string `json:"ClassName"`
	UserName string `json:"UserName	"`
	AdminType int `json:"AdminType"`
}

type ClassList struct {
	Total int `json:"Total"`
	Classes []*Class `json:"Classes"`
}
type Class struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Level int `json:"Level"`
	ParentId int `json:"ParentId"`
}

//通过用户id，查询有关联的用户
func (w *WeChat)GetUserRelations(orgId string, data RequestData) (*RelationDataList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetUserRelations", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}
	list := new(RelationDataList)
	err = response.DecodeData(list)
	return list, err
}

//通过教师获取任课班级
func (w *WeChat)GetTeacherClass(orgId string, data RequestData) (*TeacherClassList, error) {
	return w.getClass(orgId,http.MethodGet,"GetTeacherClass", data)
}

//通过班级获取老师
func (w *WeChat)GetClassTeachers(orgId string, data RequestData) (*TeacherClassList, error) {
	return w.getClass(orgId,http.MethodGet,"GetClassTeachers", data)
}
//获取班级列表，通过班级获取老师
func (w *WeChat)GetClass(orgId string, data RequestData) (*ClassList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetClass", ORG_USER_API, data)
	if err!=nil {
		return nil, err
	}
	res := new(ClassList)
	err = response.DecodeData(res)
	return res, err
}

func (w *WeChat)getClass(orgId,method, action string, data RequestData) (*TeacherClassList, error) {
	response, err := w.AuthAndRequest(orgId, method, action, ORG_USER_API, data)
	if err != nil {
		return nil, err
	}
	res := new(TeacherClassList)
	err = response.DecodeData(&res)
	return res, err
}
