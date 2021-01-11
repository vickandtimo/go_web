package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	//2.连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//3.把模型与数据库中的表对应起来
	//db.Debug().AutoMigrate(&User{})

	////4.创建
	//u1 := User{Name: "q1mi", Age: 18, Active: true}
	//db.Debug().Create(&u1)
	//u2 := User{Name: "jinzhu", Age: 20, Active: false}
	//db.Debug().Create(&u2)

	////5.查询
	//var user User
	//db.First(&user)

	////6.更新
	//user.Name = "七米"
	//user.Age = 39
	//db.Debug().Save(&user) //默认会修改所有字段
	//db.Debug().Model(&user).Update("name", "小王子")

	//让users表中所有的用户的年龄在原来的基础上+2
	db.Model(&User{}).Update("age", gorm.Expr("age+?", "2"))

}
