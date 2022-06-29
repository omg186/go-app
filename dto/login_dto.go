package dto

type LoginRequest struct {
	UserName string `json:"userName" binding:"required"` // 用户名
	PassWord string `json:"passWord" binding:"required"` // 密码
}
