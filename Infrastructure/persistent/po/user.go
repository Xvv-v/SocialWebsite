package po

import (
	"github.com/jinzhu/gorm"
	"../repository"
)
//User 对象
type User struct {
	gorm.Model
	Userid       string `gorm:"type:varchar(225);unique_index;not null;primary_key" json:"userid,omitempty"`
	Username     string `gorm:"type:varchar(220);not null" json:"username,omitempty"`
	Telephone    string `gorm:"type:varchar(12)" json:"telephone,omitempty"`
	Email        string `gorm:"type:varchar(30)" json:"email,omitempty"`
	Password     string `gorm:"type:varchar(15);not null" json:"password,omitempty"`
	Faceicon     string `gorm:"type:varchar(255);not null" json:"faceicon,omitempty"`
	Gender       string `gorm:"varchar(3);not null" json:"gender,omitempty"`
	Region       string `gorm:"varchar(50)" json:"region,omitempty"`
	Education    string `gorm:"varchar(255)" json:"education,omitempty"`
	Introduction string `gorm:"varchar(255)" json:"introduction,omitempty"`
}

//禁止使用默认表名
repository.mysqldb.SingularTable(true)