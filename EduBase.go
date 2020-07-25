package wechat

import "net/http"

type ClassRoomList struct {
	PageInfo PageInfo     `json:"PageInfo"`
	DataList []*ClassRoom `json:"DataList"`
}

type ClassRoom struct {
	BuildId   string      `json:"BuildId"`
	BuildName string      `json:"BuildName"`
	RoomInfo  []*RoomInfo `json:"RoomInfo"`
}

type RoomInfo struct {
	RoomId   string `json:"RoomId"`
	RoomName string `json:"RoomName"`
	SeatNum  string `json:"SeatNum"`
}

type YearList struct {
	PageInfo PageInfo `json:"PageInfo"`
	DataList []*Year  `json:"DataList"`
}

type Year struct {
	YearId     string   `json:"YearId"`
	Operator   string   `json:"Operator"`
	CreateTime string   `json:"CreateTime"`
	StartYear  string   `json:"StartYear"`
	EndYear    string   `json:"EndYear"`
	Opts       []string `json:"Opts"`
	Terms      []Term   `json:"Terms"`
}

type Term struct {
	TermId   string   `json:"TermId"`
	Name     string   `json:"Name"`
	TermDate []string `json:"TermDate"`
}
type TermList struct {
	TermList []*TermDetail `json:"TermList"`
}

type TermDetail struct {
	TermId     string `json:"TermId"`
	TermName   string `json:"TermName"`
	StartDate  string `json:"StartDate"`
	EndDate    string `json:"EndDate"`
	CreateDate string `json:"CreateDate"`
	YearId     int    `json:"YearId"`
	Label      string `json:"Label"`
	ClassId    int    `json:"ClassId"`
}

type CourseList struct {
	PageInfo PageInfo  `json:"PageInfo"`
	DataList []*Course `json:"DataList"`
}
type Course struct {
	Id         int      `json:"Id" json:"CourseId"`
	Name       string   `json:"Name"`
	Color      string   `json:"Color"`
	Operator   string   `json:"Operator"`
	CreateTime string   `json:"CreateTime"`
	Opts       []string `json:"Opts"` //["edit", "delete"] //操作(给前端用)
}

type CourseRelationList struct {
	CourseList []*CourseRelation `json:"CourseList"`
}
type CourseRelation struct {
	CourseIds []int `json:"CourseIds"`
	GradeId   int   `json:"GradeId"`
	Type      []int `json:"Type"`
}

type AdministrativeTeachList struct {
	List []*AdministrativeTeach `json:"List"`
}

//班级任课老师
type AdministrativeTeach struct {
	CourseId   int    `json:"CourseId"`
	TeacherId  []int  `json:"TeacherId"`
	CourseName string `json:"CourseName"`
}

type LessonList struct {
	DataList []*Lesson `json:"DataList"`
}

type Lesson struct {
	LessonId   int    `json:"LessonId"`
	LessonName string `json:"LessonName"`
	SummerTime string `json:"SummerTime"`
	WinterTime string `json:"WinterTime"`
}

// 老师课表
type TeacherScheduleList struct {
	WeekInfo WeekInfo           `json:"WeekInfo"`
	DataList []*TeacherSchedule `json:"DataList"`
}

type WeekInfo struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}
type TeacherSchedule struct {
	WeekDay int             `json:"WeekDay"`
	List    []*ScheduleList `json:"List"`
}
type ScheduleList struct {
	ClassId    int    `json:"ClassId"`
	Type       int    `json:"Type"`
	CourseId   int    `json:"CourseId"`
	WeekStart  int    `json:"WeekStart"`
	WeekEnd    int    `json:"WeekEnd"`
	LessonTime string `json:"LessonTime"`
	WeekPeriod string `json:"WeekPeriod"`
	CourseName string `json:"CourseName"`
	Teacher    string `json:"Teacher"`
	BuildName  string `json:"BuildName"`
	RoomName   string `json:"RoomName"`
}

// 班级课表
type ClassScheduleList struct {
	LessonId int              `json:"LessonId"`
	List     []*ClassSchedule `json:"List"`
}
type ClassSchedule struct {
	WeekDay int                    `json:"WeekDay"`
	List    []*ClassScheduleDetail `json:"List"`
}
type ClassScheduleDetail struct {
	ScheduleId int    `json:"ScheduleId"`
	WeekStart  int    `json:"WeekStart"`
	WeekEnd    int    `json:"WeekEnd"`
	WeekPeriod string `json:"WeekPeriod"`
	CourseName string `json:"CourseName"`
	Teacher    string `json:"Teacher"`
	BuildName  string `json:"BuildName"`
	RoomName   string `json:"RoomName"`
}

type StudentScheduleData struct {
	WeekInfo WeekInfo           `json:"WeekInfo"`
	DataList []*StudentSchedule `json:"DataList"`
}

type StudentSchedule struct {
	WeekDay int                      `json:"WeekDay"`
	List    []*StudentScheduleDetail `json:"List"`
}
type StudentScheduleDetail struct {
	ClassId    int    `json:"ClassId"`
	CourseId   int    `json:"CourseId"`
	TeacherId  int    `json:"TeacherId"`
	Type       int    `json:"Type"`
	WeekStart  int    `json:"WeekStart"`
	WeekEnd    int    `json:"WeekEnd"`
	LessonTime string `json:"LessonTime"`
	WeekPeriod string `json:"WeekPeriod"`
	CourseName string `json:"CourseName"`
	Teacher    string `json:"Teacher"`
	BuildName  string `json:"BuildName"`
	RoomName   string `json:"RoomName"`
}

type CourseSchedule struct {
	ClassId    int    `json:"ClassId"`
	Course     Course `json:"Course"`
	Teacher    []*Teacher
	Room       CourseRoom      `json:"Room"`
	WeekStart  int             `json:"WeekStart"`
	WeekEnd    int             `json:"WeekEnd"`
	WeekPeriod int             `json:"WeekPeriod"`
	Lessons    []*CourseLesson `json:"Lessons"`
}

type CourseLesson struct {
	WeekDay  int    `json:"WeekDay"`
	LessonId string `json:"LessonId"`
}

type CourseRoom struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

//获取教学楼教室详细信息列表
func (w *WeChat) GetClassroomInfoList(orgId string, data RequestData) (*ClassRoomList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetClassroomInfoList", EDU_API, data)
	if err != nil {
		return nil, err
	}
	res := new(ClassRoomList)
	err = response.DecodeData(res)
	return res, err
}

//获取学年学期列表
func (w *WeChat) GetYearList(orgId string, data RequestData) (*YearList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetYearList", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(YearList)
	err = response.DecodeData(list)
	return list, err
}

//获取学年详情
func (w *WeChat) GetYearDetail(orgId string, data RequestData) (*Year, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetYearDetail", EDU_API, data)
	if err != nil {
		return nil, err
	}

	year := new(Year)
	err = response.DecodeData(year)
	return year, err
}

//获取学期的总周数
func (w *WeChat) GetTermWeekTotal(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodGet, "GetTermWeekTotal", EDU_API, data)
}

//获取当前时间对应学期
func (w *WeChat) GetCurrentTerm(orgId string, data RequestData) (*TermList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetCurrentTerm", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(TermList)
	err = response.DecodeData(list)
	return list, err
}

//获取当前所在学期周
func (w *WeChat) GetCurrentWeek(orgId string, data RequestData) (*Response, error) {
	return w.AuthAndRequest(orgId, http.MethodGet, "GetCurrentWeek", EDU_API, data)
}

//获取学年学期列表
func (w *WeChat) GetCourseList(orgId string, data RequestData) (*CourseList, error) {
	return w.courseList(orgId, http.MethodGet, "GetCourseList", data)
}

//获取学科年级关系列表
func (w *WeChat) GetGradeCourseRelation(orgId string, data RequestData) (*CourseRelationList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetGradeCourseRelation", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(CourseRelationList)
	err = response.DecodeData(list)
	return list, err
}

//获取年级关联学科列表
func (w *WeChat) GetOpenCourseList(orgId string, data RequestData) (*CourseList, error) {
	return w.courseList(orgId, http.MethodGet, "GetOpenCourseList", data)
}

func (w *WeChat) courseList(orgId, method, action string, data RequestData) (*CourseList, error) {
	response, err := w.AuthAndRequest(orgId, method, action, EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(CourseList)
	err = response.DecodeData(list)
	return list, err
}

//获取班级任课列表
func (w *WeChat) GetAdministrativeTeachList(orgId string, data RequestData) (*AdministrativeTeachList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetAdministrativeTeachList", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(AdministrativeTeachList)
	err = response.DecodeData(list)
	return list, err
}

// 获取班级作息节次
func (w *WeChat) GetClassRestLesson(orgId string, data RequestData) (*LessonList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetClassRestLesson", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(LessonList)
	err = response.DecodeData(list)
	return list, err
}

//获取老师课表
func (w *WeChat) GetTeacherSchedule(orgId string, data RequestData) (*TeacherScheduleList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetTeacherSchedule", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(TeacherScheduleList)
	err = response.DecodeData(list)
	return list, err
}

//获取班级课表
func (w *WeChat) GetClassSchedule(orgId string, data RequestData) (*ClassScheduleList, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetTeacherSchedule", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(ClassScheduleList)
	err = response.DecodeData(list)
	return list, err
}

//获取学生课表
func (w *WeChat) GetStudentSchedule(orgId string, data RequestData) (*StudentScheduleData, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetStudentSchedule", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(StudentScheduleData)
	err = response.DecodeData(list)
	return list, err
}

//获取排课详情
func (w *WeChat) GetCourseSchedule(orgId string, data RequestData) (*CourseSchedule, error) {
	response, err := w.AuthAndRequest(orgId, http.MethodGet, "GetCourseSchedule", EDU_API, data)
	if err != nil {
		return nil, err
	}
	list := new(CourseSchedule)
	err = response.DecodeData(list)
	return list, err
}
