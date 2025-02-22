package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

// Order represents the order structure from order-service
type Order struct {
	ID            string    `json:"id"`
	UserEmail     string    `json:"user_email"`
	Product       string    `json:"product"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	PaymentStatus string    `json:"payment_status"`
	EmailStatus   string    `json:"email_status"`
}

func main() {
	// Create a router
	router := http.NewServeMux()

	// Notification handler
	router.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		log.Println("🔍 Checking payment status...")

		if err := checkPaymentStatus(); err != nil {
			http.Error(w, "Failed to check payment status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Notification process completed ✅"})
	})

	// Start the notification service
	fmt.Println("📩 Notification Service running on port 8084...")
	log.Fatal(http.ListenAndServe(":8084", router))
}

// checkPaymentStatus fetches orders and updates email status if payment is completed
func checkPaymentStatus() error {
	resp, err := http.Get("http://order-service:8081/orders")
	if err != nil {
		log.Println("Error fetching orders 😭:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to fetch orders ❌")
		return fmt.Errorf("failed to fetch orders")
	}

	var orders []Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		log.Println("Error decoding data 😵‍💫")
		return err
	}

	for _, order := range orders {
		if order.PaymentStatus == "completed" && order.EmailStatus != "sent" {
			// Update email status in order-service
			if err := updateEmailStatus(order.ID); err != nil {
				log.Printf("❌ Failed to update email status for Order ID: %s: %s\n", order.ID, err)
				continue
			}

			// Send email to user
			log.Printf("✅ Payment successful for Order ID: %s, User: %s, Product: %s\n Price %f\n", order.ID, order.UserEmail, order.Product, order.Price)
			if err := sendEmail(order.UserEmail, order.ID, order.Product, order.Price); err != nil {
				log.Printf("❌ Failed to send email: %s\n", err)
			}
		}
	}
	return nil
}

// updateEmailStatus updates the email status in the order-service
func updateEmailStatus(orderID string) error {
	requestBody, err := json.Marshal(map[string]string{"id": orderID})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://order-service:8081/update-email", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update email status, status code: %d", resp.StatusCode)
	}

	log.Printf("📨 Email status updated for Order ID: %s\n", orderID)
	return nil
}

// sendEmail simulates sending an email
func sendEmail(userEmail, orderID, product string, price float64) error {
	// Simulate email sending log
	log.Printf("📩 Sending email to %s about Order ID: %s for Product: %s\n", userEmail, orderID, product)

	// Simulate delay (optional)
	time.Sleep(2 * time.Second)

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"                  // Change if using another provider
	smtpPort := "587"                             // TLS port
	senderEmail := "arpitsrivastava529@gmail.com" // Your email (use environment variable)
	senderPass := "bxxp vrkd fhku tqgz"           // App password (use environment variable)

	// Email message
	subject := "Your Order Confirmation"
	body := fmt.Sprintf(
		"Hello,\n\nYour order (ID: %s) for %s has been confirmed. Total Price is %f\nThank you for shopping with us!\n\nBest regards,\nYour KhuPrit Team",
		orderID, product, price,
	)
	msg := "From: " + senderEmail + "\n" +
		"To: " + userEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// SMTP authentication
	auth := smtp.PlainAuth("", senderEmail, senderPass, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{userEmail}, []byte(msg))
	if err != nil {
		log.Printf("❌ Failed to send email to %s: %v\n", userEmail, err)
		return err
	}

	log.Println("✅ Email sent successfully!")
	return nil
}
