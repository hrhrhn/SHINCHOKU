package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Test_user struct {
	Id       int    `gorm:"column:id"`
	UserName string `gorm:"column:parent_name"`
}

type Test_todo struct {
	Id      int    `gorm:"column:id"`
	Content string `gorm:"column:content"`
	Done    bool   `gorm:"column:done"`
	UserId  int    `gorm:"column:user_id"`
}

// connect to the Db
func main() {
	db, err := gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		fmt.Println("--init--")
	}
	defer db.Close()

	user := Test_user{UserName: "testUser001"}
	fmt.Println(user)

	// Create a user
	db.Create(&user)
	fmt.Println(user)

	// Add a todo
	// id := nil
	todo := Test_todo{Content: "Learn Go Programming nil"}
	db.Create(&todo)

	// Get comments from a post
	// var readTodo Todo
	// err = db.Where("user_id = $1", user.Id).First(&readTodo).Error
	// if err != nil {
	// 	log.Fatalln("取得失敗", err)
	// }
	// fmt.Println(readTodo)
}
