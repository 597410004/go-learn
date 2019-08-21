package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Num  int `gorm:"AUTO_INCREMENT"`
	CreditCard CreditCard  `gorm:"save_associations:false"`
}
type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

func TestMysql(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/gorm_learn?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.CreateTable(new(User))

}
