package routes

import (
	"github.com/Arpit529stivastava/payment-services/handler"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	paymentGroup := r.Group("/pay")
	{
		paymentGroup.POST("/", handler.ProcessPayment)
		paymentGroup.GET("/:id", handler.GetPaymentStatus)
	}
}
