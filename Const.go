package wechat

type ErrCode string
type OrgType string

const (
	OFFICE OrgType = "上级单位"
	GROUP  OrgType = "总集"
	SCHOOL OrgType = "学校"
)

const (
	ORG_USER_API = "/v2/user"
	MSG_API      = "/v2/message"
)

const (
	ERROR_CODE_OK     ErrCode = "OK"
	ERROR_CODE_FAILED ErrCode = "Failed"
)
