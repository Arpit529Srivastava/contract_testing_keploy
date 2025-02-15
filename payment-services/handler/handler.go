package handler

import (
	"net/http"

	"strconv"

	"github.com/Arpit529stivastava/payment-services/services"
	"github.com/gin-gonic/gin"
)

func ProcessPayment(c *gin.Context) {
	var req struct {
		OrderID int     `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentID, err := services.ProcessPayment(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_id": paymentID, "status": "completed"})
}

func GetPaymentStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	status, err := services.GetPaymentStatus(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_status": status})
}
