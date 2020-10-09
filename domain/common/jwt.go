package common

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyClaim 储存信息
type MyClaim struct {
	Userid   string
	Username string
	Faceicon string
	jwt.StandardClaims
}

const (
	//TokenExpireDuration 过期时间
	TokenExpireDuration = time.Hour * 2
	//MySecret 密钥
	MySecret = "xjw0331"
)

//CreateAccessToken 生成token
func CreateAccessToken(id, name, icon string) string {

	//创建声明
	claim := MyClaim{
		id,
		name,
		icon,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Xvv-v",
		},
	}

	//创造签名对象
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//使用指定的签名获得完整的token
	token, err := accessToken.SignedString([]byte(MySecret))
	if err != nil {
		log.Println(err)
		return ""
	}

	return token
}

//ParaseAccessToken 解析jwt
func ParaseAccessToken(token string)  {
	
	//解析token
}