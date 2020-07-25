# 腾讯智慧校园 SDK

```go
func main()  {
    // AppId, SecretKey, ApiUrl 为必填。
    conf := &Config{
        AppId: 381232,
        SecretKey: "your_secret_key",
        ApiAddr: "oapi.campus.qq.com",
        Timeout: time.Duration(2) * time.Second
    }
    // 创建SDK实例
    client := NewWeChat(conf)
    // 设置网络异常时重试次数
    client.SetRetry(3)
    // 请求数据，按照智慧校园官方文档设置map的key和值
    var data = make(map[string]interface{})
    data["OrgUserId"] = []string{"1", "2"}
    // 获取用户信息
    info, err := client.GetUsersInfo("188776", data)
    if err != nil {
        return
    }
    // 返回的结果为可对应的结构体，以直接使用。
    fmt.Println(info.Total)
    fmt.Println(info.DataList)
}
```
