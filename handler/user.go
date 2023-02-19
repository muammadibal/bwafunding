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

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// fmt.Println(err.Error())
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Account register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Account register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwt

	formatter := user.FormatUser(userData, "tokennnnnn")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
