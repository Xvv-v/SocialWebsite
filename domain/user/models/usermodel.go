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

func getPwd(user *User) string {
	
	return user.password
}

