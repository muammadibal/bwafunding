package main

import (
	"bwafunding/auth"
	"bwafunding/campaign"
	"bwafunding/handler"
	"bwafunding/transaction"
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

	campaignRepository := campaign.AssignRepository(db)
	campaignService := campaign.AssignService(campaignRepository)
	campaignHandler := handler.AssignCampaignHandler(campaignService)

	transactionRepository := transaction.AssignRepository(db)
	transactionService := transaction.AssignService(transactionRepository, campaignRepository)
	transactionHandler := handler.AssignTransactionHandler(transactionService)
	// input := campaign.CampaignCreateInput{}
	// input.Name = "Penggalangan Dana Startup"
	// input.ShortDescription = "short"
	// input.Description = "loooonnngggg"
	// input.GoalAmount = 1000000000
	// input.Perks = "hadiah satu, dua, dan tiga"

	// inputUser, _ := userService.UserDetail(5)
	// input.User = inputUser
	// inputSave, err := campaignService.CreateCampaign(input)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println(inputSave)

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "test lagi"
	// userInput.Email = "test@example.com"
	// userInput.Occupation = "baru join"
	// userInput.Password = "password"
	// userService.RegisterUser(userInput)

	// var users []user.User
	// db.Find(&users)

	router := gin.Default()
	router.Static("/images", "./images")
	apiV1 := router.Group("/api/v1")
	apiV1.POST("/users/sign-up", userHandler.Register)
	apiV1.POST("/users/sign-in", userHandler.Login)
	apiV1.POST("/users/email_checkers", userHandler.CheckAvailabilityEmail)
	apiV1.POST("/users/upload_avatar", handler.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	apiV1.GET("/users/fetch", userHandler.FetchUser)

	apiV1.GET("/campaigns", campaignHandler.Campaigns)
	apiV1.POST("/campaigns", handler.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	apiV1.GET("/campaigns/:id", campaignHandler.CampaignsDetail)
	apiV1.PUT("/campaigns/:id", handler.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	apiV1.POST("/campaign-images", handler.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	apiV1.GET("/campaigns/:id/transactions", handler.AuthMiddleware(authService, userService), transactionHandler.CampaignTransactions)
	apiV1.GET("/transactions", handler.AuthMiddleware(authService, userService), transactionHandler.CampaignTransactionsUser)

	router.Run()
}
