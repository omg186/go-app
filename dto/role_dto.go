package dto

type RoleRes struct {
	ID       int    `json:"id"`
	RoleName string `json:"roleName"` // 角色名
	Value    string `json:"value"`    // 角色code
	Desc     string `json:"desc"`     // 描述
}
