package routes

import (
	handlers "github.com/Arpit529stivastava/notification-services/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/notify", handlers.SendNotification)
	r.GET("/status/:id", handlers.GetStatus)
	r.POST("/payment-success", handlers.PaymentSuccessEvent)
}
