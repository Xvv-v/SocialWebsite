package model

//Register 定义注册数据模型
type Register struct {
	Username     string `json:"username,omitempty"`     //用户名
	Telephone    string `json:"telephone,omitempty"`    //手机号
	Email        string `json:"email,omitempty"`        //邮箱
	Password     string `json:"password,omitempty"`     //密码
	Verification string `json:"verification,omitempty"` //验证码
	Gender       string `json:"gender,omitempty"`       //性别
	Region       string `json:"region,omitempty"`       //地区
	Education    string `json:"education,omitempty"`    //教育经历
}

//Login 定义登陆数据模型
type Login struct {
	Account  string `json:"account，omitempty"`  //账号
	Password string `json:"password，omitempty"` //密码
}
