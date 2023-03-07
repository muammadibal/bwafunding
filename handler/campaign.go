package handler

import (
	"bwafunding/campaign"
	"bwafunding/helper"
	"bwafunding/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service menentukan repository yg akan di call
// repository : pertama
// db : kedua d

type campaignHandler struct {
	service campaign.Service
}

func AssignCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CampaignCreateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Error create campaigns", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println("raw current user", c.MustGet("currentUser"))
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	campaignData, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Error create campaigns", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignData))

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	var input campaign.CampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CampaignCreateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error get detail campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	inputData.User = currentUser
	updateCampaign, err := h.service.UpdateCampaign(input, inputData)
	if err != nil {
		response := helper.APIResponse("Error update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update campaign", http.StatusOK, "success", updateCampaign)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) Campaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.Campaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CampaignsDetail(c *gin.Context) {
	// ID, _ := strconv.Atoi(input.ID)
	var input campaign.CampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignData, err := h.service.CampaignDetail(input)
	if err != nil {
		response := helper.APIResponse("Error get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := campaign.FormatCampaignDetail(campaignData)
	response := helper.APIResponse("Success get detail campaign", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	input.User = currentUser

	path := fmt.Sprintf("images/campaign/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.CreateCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success upload campaign image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
