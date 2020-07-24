# 腾讯智慧校园 SDK

```go
func main()  {
    conf := &Config{}
    // AppId, SecretKey, ApiUrl
    client := NewWeChat(600001, "your_secret_key", "oapi.campus.qq.com",conf)
    var data = make(map[string]interface{})
    data["OrgUserId"] = []string{"1", "2"}
    info, err := client.GetUsersInfo("188776", data)
    if err != nil {
        return
    }
    fmt.Println(info.Total)
    fmt.Println(info.DataList)
}
```
