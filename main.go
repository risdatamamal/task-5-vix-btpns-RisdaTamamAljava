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
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.PhoneNumber = "08202092092"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
