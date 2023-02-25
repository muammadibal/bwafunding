package handler

import (
	"bwafunding/campaign"
	"bwafunding/helper"
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

func (h *campaignHandler) Campaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.Campaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaigns)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CampaignsDetail(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	campaign, err := h.service.CampaignDetail(ID)
	if err != nil {
		response := helper.APIResponse("Error get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success get campaign", http.StatusOK, "success", campaign)

	c.JSON(http.StatusOK, response)
}
