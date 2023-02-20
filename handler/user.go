package handler

import (
	"bwafunding/helper"
	"bwafunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func AssignUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// fmt.Println(err.Error())
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userData, err := h.userService.Register(input)
	if err != nil {
		response := helper.APIResponse("Register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwt

	formatter := user.FormatUser(userData, "tokennnnnn")

	response := helper.APIResponse("Register Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userData, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userData, "tokennnnnn")

	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckAvailabilityEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := h.userService.CheckAvailabilityEmail(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Email checking failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	status := gin.H{
		"is_available": isAvailable,
	}

	metaMessage := "Email has been registered"
	if isAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", status)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Fetch user data success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
