package handlers

import (
	"log"
	"net/http"

	"github.com/Arpit529stivastava/notification-services/config"
	"github.com/Arpit529stivastava/notification-services/models"
	"github.com/gin-gonic/gin"
)

// Send Notification
func SendNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO notifications (user_id, message, status) VALUES ($1, $2, 'pending')",
		notification.UserID, notification.Message)

	if err != nil {
		log.Println("Error inserting notification:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}

// Get Notification Status
func GetStatus(c *gin.Context) {
	id := c.Param("id")

	var status string
	err := config.DB.QueryRow("SELECT status FROM notifications WHERE id = $1", id).Scan(&status)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

// Receive Payment Success Event
func PaymentSuccessEvent(c *gin.Context) {
	var payload struct {
		UserID int `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	message := "Your payment was successful!"
	_, err := config.DB.Exec("INSERT INTO notifications (user_id, message, status) VALUES ($1, $2, 'sent')",
		payload.UserID, message)

	if err != nil {
		log.Println("Error inserting payment notification:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to notify user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment notification sent"})
}
