package po

import (
	"SocialWebsite/infrastructure/repository"
	"errors"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

//UserPO 持久化对象
type UserPO struct {
	Userid    int64  `gorm:"type:int;primaryKey;unique;not null;column:user_id"`
	Username  string `gorm:"type:varchar;size:255;not null;column:username"` //用户名
	Telephone string `gorm:"type:varchar;size:255;column:telephone"`         //手机号
	Email     string `gorm:"type:varchar;size:255;column:email"`             //邮箱
	Password  string `gorm:"type:varchar;size:255;not null;column:password"` //密码
	Faceicon  string `gorm:"type:varchar;size:255;not null;column:faceicon"` //头像
	Gender    string `gorm:"type:varchar;size:255;column:gender"`            //性别
	Region    string `gorm:"type:varchar;size:255;column:region"`            //地区
	Education string `gorm:"type:varchar;size:255;column:education"`         //教育经历
}

//TableName 将 User 的表名设置为 `profiles`
func (UserPO) TableName() string {
	return "user"
}

//NewUserPO 创建新po对象
func NewUserPO(id int64, name, tele, email, pwd, face, gender, reg, edu string) *UserPO {

	return &UserPO{
		Userid:    id,
		Username:  name,
		Telephone: tele,
		Email:     email,
		Password:  pwd,
		Faceicon:  face,
		Gender:    gender,
		Region:    reg,
		Education: edu,
	}
}

//Insert 插入
func Insert(u *UserPO) error {

	return repository.Mysqldb.Create(&u).Error
}

//SelectRow 单行查询
func SelectRow(sql string, value ...interface{}) (*UserPO, bool) {

	user := &UserPO{}
	err := repository.Mysqldb.Where(sql, value).First(&user).Error
	return user, errors.Is(err, gorm.ErrRecordNotFound)
}

//SelectMulit 多行查询
func SelectMulit(sql string, value ...interface{}) ([]UserPO, error) {

	var result []UserPO
	if err := repository.Mysqldb.Where(sql, value...).Find(&result).Error; err == nil {
		//当find找不到记录时error返回nil
		return nil, errors.New("找不到记录")
	}
	return result, nil
}

//Update 更新选定字段
func Update(sql, column, value1, value2 string) error {

	return repository.Mysqldb.Model(&UserPO{}).Where(sql, value1).Update(column, value2).Error
}

//Delete 删除
func Delete(sql, value string) error {

	return repository.Mysqldb.Where(sql, value).Delete(&UserPO{}).Error
}

//Set redis操作
func Set(key int64, value string) error {
	
	//过期时间：time.Hour * 2
	_,err:=repository.Redisdb.Do("SET",key,value,"EX",2)
	return err
}

//Get redis操作
func Get(key int64) error {
	
	_,err:=repository.Redisdb.Do("GET",key)
	if err!=nil{
		if err==redis.Nil{
			return errors.New("找不到纪录")
		}
	}
	return err
}
