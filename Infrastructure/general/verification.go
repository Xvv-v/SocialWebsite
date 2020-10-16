package general

import (
	"fmt"
	"math/rand"
	"time"
)

//生成验证码
func getCode() string {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

//SendSMSCode 发送短信验证
func SendSMSCode() {

}

//SendEmailCode 发送邮箱验证
func SendEmailCode() {

}
