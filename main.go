package main

import (
	"log"
	"task-5-vix-btpns-RisdaTamamAljava/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_starcamp_academy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Test simpan",
		Email: "test@gmail.com",
		Password: "test123",
		Roles: "USER",
	}

	userRepository.Save(user)
}
