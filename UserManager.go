package wechat

type UserAddUpdateResponse struct {
	Success []*SuccessIdx `json:"Success"`
	Fails   []*FailedIdx  `json:"fails"`
}
type SuccessIdx struct {
	Idx       int    `json:"Idx"`
	OrgUserId string `json:"OrgUserId"`
}
type FailedIdx struct {
	Idx     int    `json:"Idx"`
	Message string `json:"Message"`
}

type Teacher struct {
	Name     string `json:"Name"`
	Phone    string `json:"Phone"`
	Gender   int    `json:"Gender"`   //1：男 2：女
	UserNo   string `json:"UserNo"`   //教工号
	JoinDate string `json:"JoinDate"` // 2019-09-10
	// example: [{"Name": "教务处","Code":"jgabc"},{"Name": "校长室","Code":"jg123"}]
	Departments []*TeacherDepartment `json:"Departments"`
}

type TeacherDepartment struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type DeleteTeachersRes struct {
	OrgUserIds []string    `json:"OrgUserIds"`
	Fails      []FailedIdx `json:"Fails"`
}

type UserDeleteRes struct {
	Fails      []FailedIdx `json:"Fails"`
}

type OpenStudent struct {
	Name string `json:"Name"`
	UserNo string `json:"UserNo	"`
	JoinDate string `json:"JoinDate"`
	Sex int `json:"Sex"` //性别 1:男 2:女
	Departments []*StudentDepartment `json:"Departments"`
}
type StudentDepartment struct {
	Id int `json:"Id"` //架构id
}

type OpenParent struct {
	StudentId string `json:"StudentId"`
	Relation string `json:"Relation"`
	Name string `json:"Name"`
	Phone string `json:"Phone"`
	Work string `json:"Work"`
}

//添加教师, data数据示例： data["Teachers"] = []*Teachers
func (w *WeChat) AddTeachers(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId, data, "AddTeachers")
}

//修改教师, data数据示例： data["Teachers"] = []*Teachers
func (w *WeChat) UpdateTeachers(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId, data, "UpdateTeachers")
}

//添加临时成员  data数据示例： data["Teachers"] = []*Teachers
func (w *WeChat) AddTemporarys(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId, data, "AddTemporarys")
}

//修改临时成员 data数据示例： data["OrgUserIds"] = []*Teachers
func (w *WeChat) UpdateTemporarys(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId, data, "UpdateTemporarys")
}

//添加学生信息，支持批量添加 data["OpenStudent"] = []*OpenStudent
func (w *WeChat)AddStudents(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId,data,"AddStudents")
}

//修改学生信息，支持批量修改 data["OpenStudent"] = []*OpenStudent
func (w *WeChat)UpdateStudents(orgId string, data RequestData) (*UserAddUpdateResponse, error) {
	return w.addAndUpdate(orgId,data,"UpdateStudents")
}

// 添加家长信息，支持批量添加 data["OpenParent"] = []*OpenParent
func (w *WeChat)AddParents(orgId string, data RequestData) (*UserAddUpdateResponse, error)  {
	return w.addAndUpdate(orgId,data, "AddParents")
}

// 修改家长信息，支持批量修改 data["OpenParent"] = []*OpenParent
func (w *WeChat)UpdateParents(orgId string, data RequestData) (*UserAddUpdateResponse, error)  {
	return w.addAndUpdate(orgId,data, "UpdateParents")
}

//删除教师 data["Teachers"] = []*Teachers
func (w *WeChat) DeleteTeachers(orgId string, data RequestData) (*DeleteTeachersRes, error) {
	response, err := w.authAndRequest(orgId, "POST", "DeleteTeachers", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var res DeleteTeachersRes
	if err := response.DecodeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

//批量删除学生 data["OpenStudent"] = []*OrgUserId{"1", "2"}
func (w *WeChat) DeleteStudents(orgId string, data RequestData) (*UserDeleteRes, error)  {
	response, err := w.authAndRequest(orgId, "POST", "DeleteStudents", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var res UserDeleteRes
	if err := response.DecodeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

//批量删除家长 data["OpenParent"] = []*OrgUserId{"1", "2"}
func (w *WeChat) DeleteParents(orgId string, data RequestData) (*UserDeleteRes, error)  {
	response, err := w.authAndRequest(orgId, "POST", "DeleteParents", ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var res UserDeleteRes
	if err := response.DecodeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (w *WeChat) addAndUpdate(orgId string, data RequestData, action string) (*UserAddUpdateResponse, error) {
	response, err := w.authAndRequest(orgId, "POST", action, ORG_USER_API, data)
	if err != nil {
		return nil, err
	}

	var res UserAddUpdateResponse
	if err := response.DecodeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
