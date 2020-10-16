package model

//UserVO 定义注册数据模型
type UserVO struct {
	Username  string `json:"username,omitempty"`  //用户名
	Telephone string `json:"telephone,omitempty"` //手机号
	Email     string `json:"email,omitempty"`     //邮箱
	Password  string `json:"password,omitempty"`  //密码
}

//LoginVO 定义登陆数据模型
type LoginVO struct {
	Telephone string `json:"telephone" binding:"telephone"` //手机号
	Email     string `json:"email" binding:"email"`         //邮箱
	Password  string `json:"password" binding:"password"`   //密码
}
