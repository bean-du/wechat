# 腾讯智慧校园 SDK

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
    
	// 下面是POST请求传参方法：

}
```
