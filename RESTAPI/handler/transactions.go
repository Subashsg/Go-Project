package handler

import (
	user "User/User"
	"User/helper"
	"User/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {

	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {

	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)

	if err != nil {

		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	usertrans, err := h.service.GetTransactionsByUserID(userID)

	if err != nil {

		responce := helper.APIResponse("Failed to get the details", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, responce)
	}

	responce := helper.APIResponse("User Transactions", http.StatusOK, "success", transaction.FormatUserTransactions(usertrans))
	c.JSON(http.StatusOK, responce)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {

	var input transaction.CreateTransactionInput

	err := c.ShouldBind(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	trans, err := h.service.CreateTransaction(input)

	if err != nil {

		responce := helper.APIResponse("Failed to get the details", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, responce)
	}

	responce := helper.APIResponse("Success to create transaction", http.StatusOK, "success", transaction.FormatTransaction(trans))
	c.JSON(http.StatusOK, responce)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {

	var input transaction.TransactionNotificationInput

	err := c.ShouldBind(&input)

	if err != nil {

		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return

	}

	err = h.service.ProcessPayment(input)

	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	c.JSON(http.StatusOK, input)
}
