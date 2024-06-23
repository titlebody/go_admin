package service

// LoginRequest 接收登录参数
type LoginRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// LoginSum 登录成功后token的结构体
type LoginSum struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
