package socialwebsite

import "errors"

var (
	//ErrRecordNotFound 找不到记录
	ErrRecordNotFound = errors.New("没有该记录")
	//ErrAccountNotSignIn 账号还未注册
	ErrAccountNotSignIn = errors.New("该账号还没有注册")
)
