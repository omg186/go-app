package dto

type UserRes struct {
	Id       int         `json:"id"`
	UserName string      `json:"userName"` // 登录名
	Password string      `json:"-"`        // 密码
	RealName string      `json:"realName"` // 真实姓名
	Avatar   string      `json:"avatr"`    // 头像
	Desc     string      `json:"desc"`     // 描述
	HomePath string      `json:"homePath"` // 主页路径
	RoleList interface{} `json:"roleList"` // 角色
}
