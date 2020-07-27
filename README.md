# 腾讯智慧校园 SDK



### 初始化SDK实例

```go
func main()  {
    // AppId, SecretKey, ApiUrl 为必填。
    conf := &Config{
        AppId: 381232,
        SecretKey: "your_secret_key",
        ApiAddr: "oapi.campus.qq.com",
        // 设置请求超时时间，此设置必填
        Timeout: time.Duration(2) * time.Second
    }
    // 创建SDK实例
    client := NewWeChat(conf)
    // 设置网络异常时重试次数，默认为2次
    client.SetRetry(3)
}
```



### GET 方式没参数示例

```go
info, err := c.GetOrgInfo(orgId, nil)
if err != nil {
	return
}
fmt.Println(info)
```



### GET 方式有参数示例



```go
data := wechat.RequestData{"DepartmentType": 0, "PageIndex": 1, "PageSize": 10}
res, err := c.GetDepartmentList(orgId, data)
if err != nil {
	log.Println(err)
}

for _, v := range res.Departments {
	fmt.Println(v)
}
fmt.Println(res)
    
```



### POST 方式传参示例

```go
var students []*wechat.OpenStudent
dpt := &wechat.StudentDepartment{Id: 1}
student := &wechat.OpenStudent{
	Name: "张三",
	UserNo: "zhangsan",
	JoinDate: "2020-07-20",
	Sex: 1,
	Departments: dpt,
}
students = append(students, student)
data := wechat.RequestData{"OpenStudent": students}

res, err := c.AddStudents(orgId, data)
if err != nil  {
	log.Println(err)
	return
}
log.Println(res)
```
