package models

//User 用户
type User struct {
	userid string
	username string 
	telephone string
	email string
	password string
	faceicon string 
	gender string
	region string
	education string
}

//获得密码
func getPwd(user *User) string {
	
	return user.password
}

//获得手机号
func getTelephone(user *User) string {
	
	return user.telephone
}

//获得邮箱
func getEmail() string {
	return ""
}

//Register 注册
func Register()  {
	
	//数据没有问题，检查该账号是否已存在

	
}

//Login 登录
func Login(account,pwd string)  {
	
	//查找数据库中是否有该账号信息
}

//Logout 登出
func Logout()  {
	
}