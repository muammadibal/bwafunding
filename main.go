package main

import (
	"bwafunding/auth"
	"bwafunding/handler"
	"bwafunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3307)/bwafunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// db.AutoMigrate(&user.User{})
	// db.AutoMigrate(&campaign.Campaign{})
	// db.AutoMigrate(&campaign.CampaignImage{})
	// db.AutoMigrate(&transaction.Transaction{})

	authService := auth.AssignService()

	userRepository := user.AssignRepository(db)
	userService := user.AssignService(userRepository)
	userHandler := handler.AssignUserHandler(userService, authService)

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "test lagi"
	// userInput.Email = "test@example.com"
	// userInput.Occupation = "baru join"
	// userInput.Password = "password"
	// userService.RegisterUser(userInput)

	// var users []user.User
	// db.Find(&users)

	// fmt.Println(users)
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/users/sign-up", userHandler.Register)
	apiV1.POST("/users/sign-in", userHandler.Login)
	apiV1.POST("/users/email_checkers", userHandler.CheckAvailabilityEmail)
	apiV1.POST("/users/upload_avatar", handler.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	apiV1.GET("/users/fetch", userHandler.FetchUser)

	router.Run()
}
