package model

import (
	"FrontAccess/src/service"
	"time"
)

type User struct {
	ID   int64 				`gorm:""column:id,primary_key"`
	username string			`gorm:""column:username"`
	age int
	birthday time.Time
	sex bool
	address string
}

//设置表名，默认是结构体的名的复数形式
func (User) TableName() string {
	return "user"
}

func GetUserByName(userName string) (user []User, err error) {
	var auser []User
	//ret := dao.MysqlClient.Table("user").Select("id").Where("username = ?", userName).Take(&auser)

	ret:= service.MysqlClient.Raw("select id from user where username = ?",userName).Scan(&auser)
	return auser, ret.Error
}
