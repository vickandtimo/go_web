package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1.定义模型
//type Model struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time `sql:"index"`
//}
type User struct {
	gorm.Model //ID CreatedAt UpdatedAt DeletedAt
	Name       string
	Age        int64
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	////3.创建
	//u1 := User{Name: "CBiBank", Age: 3}
	//db.Create(&u1)
	//u2 := User{Name: "zhaoyue", Age: 30}
	//db.Create(&u2)

	//4.查询
	//var user User //声明模型结构体类型变量user
	//db.First(&user)
	//fmt.Printf("user:%#v\n", user)

	//var users []User
	//db.Debug().Find(&users)
	//fmt.Printf("users:%#v\n", users)

	//FirstOrInit
	var user User
	db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "提莫"})
	fmt.Printf("user:%#v\n", user)
}
