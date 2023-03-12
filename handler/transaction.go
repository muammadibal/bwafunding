package handler

import (
	"bwafunding/helper"
	"bwafunding/transaction"
	"bwafunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func AssignTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CampaignTransactions(c *gin.Context) {
	var input transaction.TransactionDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.TransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CampaignTransactionsUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User) // .(user.User) artinya untuk merubah returnan menjadi tipe user
	userID := currentUser.ID

	transactions, err := h.service.TransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign transactions users", http.StatusOK, "success", transaction.FormatCampaignTransactionUsers(transactions))
	c.JSON(http.StatusOK, response)
}
