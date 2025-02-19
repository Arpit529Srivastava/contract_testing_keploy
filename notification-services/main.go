package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
)

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

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		checkPaymentStatus()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "success"})
	})
	fmt.Println("Notification service is running on localhost:8084 🚀")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func checkPaymentStatus() {
	resp, err := http.Get("http://localhost:8081/orders")
	if err != nil {
		log.Println("Error fetching orders 😭")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to fetch orders ❌")
		return
	}

	var orders []Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		log.Println("Error decoding data 😵‍💫")
		return
	}

	for _, order := range orders {
		if order.PaymentStatus == "completed" && order.EmailStatus != "sent" {
			if err := updateEmailStatus(order.ID); err != nil {
				log.Printf("❌ Failed to update email status for Order ID: %s: %s\n", order.ID, err)
				continue
			}

			log.Printf("✅ Payment successful for Order ID: %s, User: %s, Product: %s\n", order.ID, order.UserEmail, order.Product)
			if err := sendEmail(order.UserEmail, order.ID, order.Product); err != nil {
				log.Printf("❌ Failed to send email: %s\n", err)
			}
		}
	}
}

func updateEmailStatus(orderID string) error {
	requestBody, err := json.Marshal(map[string]string{"id": orderID, "email_status": "sent"})
	if err != nil {
		return fmt.Errorf("error encoding email status request: %s", err)
	}

	resp, err := http.Post("http://localhost:8081/orders/update-email", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error contacting order-service: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update email status, status code: %d", resp.StatusCode)
	}

	return nil
}

func sendEmail(to, orderID, product string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	if smtpUsername == "" || smtpPassword == "" {
		return fmt.Errorf("SMTP credentials are missing! Set SMTP_USERNAME and SMTP_PASSWORD in .env file")
	}

	subject := "Your payment for your order was successful 😊"
	body := fmt.Sprintf("Dear customer,\nThank you for trusting Keploy! Your order for Order ID: %s (Product: %s) has been successfully processed.\n\nThank you for shopping with us!\nVisit us again.", orderID, product)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUsername, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending the email: %s", err)
	}

	return nil
}