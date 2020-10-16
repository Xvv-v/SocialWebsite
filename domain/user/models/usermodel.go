package models

import (
	"SocialWebsite/infrastructure/general"
	"SocialWebsite/infrastructure/po"
	"errors"
	"fmt"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/golang/glog"
)

//User 用户
type User struct {
	userid    int64
	username  string
	telephone string
	email     string
	password  string
	faceicon  string
	gender    string
	region    string
	education string
}

//Token 令牌
type Token struct {
	AccessToken  string
	Secert       string
}

//NewUserDO 创建user do对象
func NewUserDO(telephone, email, password string) *User {
	return &User{
		userid:    0,
		username:  "",
		telephone: telephone,
		email:     email,
		password:  password,
		faceicon:  "",
		gender:    "",
		region:    "",
		education: "",
	}
}

//UserToDO 将vo对象转化为do对象
// func UserToDO(u *model.User) *User {
// 	return &User{
// 		userid:    "",
// 		username:  u.Username,
// 		telephone: u.Telephone,
// 		email:     u.Email,
// 		password:  u.Password,
// 		faceicon:  "",
// 		gender:    u.Gender,
// 		region:    u.Region,
// 		education: u.Education,
// 	}
// }

//将PO数据转换成DO数据
// func userToDO(u *po.UserPO) *User {
// 	return &User{
// 		userid:    u.Userid,
// 		username:  u.Username,
// 		telephone: u.Telephone,
// 		email:     u.Email,
// 		password:  u.Password,
// 		faceicon:  u.Faceicon,
// 		gender:    u.Gender,
// 		region:    u.Region,
// 		education: u.Education,
// 	}
// }

//Register 注册
func (user *User) Register() error {

	//数据没有问题，检查该账号是否已存在
	var res bool
	if user.telephone != "" {
		//电话不为空
		_, res = po.SelectRow("telephone=?", user.telephone)
	} else {
		_, res = po.SelectRow("email=?", user.email)
	}
	if res != true {
		//没有注册过，可以注册
		//生成userid
		s, err := snowflake.NewSnowflake(int64(0), int64(0))
		if err != nil {
			glog.Error(err)
		}
		id := s.NextVal()
		fmt.Println(id)
		//将信息存入数据库
		u := po.NewUserPO(id, user.username, user.telephone, user.email, user.password,
			user.faceicon, user.gender, user.region, user.education)
		if err := po.Insert(u); err != nil {
			return nil
		}
		return errors.New("注册错误，请重新注册")
	}
	return errors.New("该账号已注册")
}

//Login 登录
func (user *User) Login(mark bool) (*Token, error) {

	var u *po.UserPO
	var ok bool
	var token string
	if mark == true {
		u, ok = po.SelectRow("telephone=?", user.telephone)
	} else {
		u, ok = po.SelectRow("email=?", user.email)
	}
	if !ok {
		//该账号存在
		//判断账号密码是否正确
		if u.Telephone == user.telephone && u.Password == user.password {
			//账号密码匹配成功,生成token
			token = general.CreateAccessToken(u.Userid, u.Username, u.Faceicon)
			//将生成的token存入redis
			po.Set(u.Userid, token)
		} else {
			return nil, errors.New("密码错误")
		}
	} else {
		//没有找到记录
		return nil, errors.New("该账号未注册")
	}
	//登陆成功，返回令牌
	return &Token{
		AccessToken: token,
		Secert:      general.MySecret,
	}, nil
}

//Logout 登出
func (user *User) Logout() {

}
