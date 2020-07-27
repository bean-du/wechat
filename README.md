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
  // 接口为GET方式，没有请求参数，参数可以传nil
    info, err := client.GetOrgInfo(&quot;188776&quot;,nil)
    if err != nil {
        return
    }
   
    // 下面为GET方法带参数的接口请求示例：参数为 RequestData 类型，实际是map
    data := wechat.RequestData{&quot;DepartmentType&quot;: 0, &quot;PageIndex&quot;: 1, &quot;PageSize&quot;: 10}
	res, err := c.GetDepartmentList(orgId, data)
	if err != nil {
		log.Println(err)
	}
```



### GET 方式有参数示例



```go
  // 次接口为GET方式，没有请求参数，参数可以传nil
    info, err := client.GetOrgInfo("188776",nil)
    if err != nil {
        return
    }
   
    // 下面为GET方法带参数的接口请求示例：参数为 RequestData 类型，实际是map
    data := wechat.RequestData{"DepartmentType": 0, "PageIndex": 1, "PageSize": 10}
	res, err := c.GetDepartmentList(orgId, data)
	if err != nil {
		log.Println(err)
	}
    
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
