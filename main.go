package main

import (
	"log"
	"task-5-vix-btpns-RisdaTamamAljava/handler"
	"task-5-vix-btpns-RisdaTamamAljava/user"

	"github.com/gin-gonic/gin"
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

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	// apiUsers := router.Group("/users")

	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.POST("/users/email_checkers", userHandler.CheckEmailAvailability)

	router.Run()
}
