package global

type LoginForm struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}
