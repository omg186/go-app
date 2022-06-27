package loginDto

type LoginRequest struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord"`
}
