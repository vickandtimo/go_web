package main

//gorm demo1

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	////创建数据行
	//u1 := UserInfo{ID: 1, Name: "CBibank", Gender: "男", Hobby: "滑雪"}
	//db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u) //查询表中的第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)

	//更新
	db.Model(&u).Update("hobby", "摩托车 ")

	//删除
	db.Delete(&u)
}
