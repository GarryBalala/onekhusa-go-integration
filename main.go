package main

import (
	"fmt"
	"log"
	"net/http"
	"onekhusa-go/app/models"
	"onekhusa-go/app/services"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
)

var ticketTracker = make(map[string]string)

func main() {
	godotenv.Load()
	service := &services.OneKhusaService{}

	// 1. Setup Socket.io
	server := socketio.NewServer(nil)
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	// 2. Setup Gin
	r := gin.Default()

	// 3. Static Files & Root Redirect
	r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/public/index.html")
	})

	// 4. API: Initiate Checkout
	r.POST("/api/Tickets/buy/:eventId", func(c *gin.Context) {
		reference := fmt.Sprintf("OT-GO%d", time.Now().Unix()%100000000)

		data, err := service.InitiateHostedCheckout(2500.00, reference)
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": err.Error()})
			return
		}

		ticketTracker[reference] = "Pending"

		redirectURL := fmt.Sprintf("https://checkout.onekhusa.com/requestToPay/initiate?ptid=%s", data.PaymentTransactionID)

		c.JSON(200, gin.H{
			"status":      "success",
			"redirectUrl": redirectURL,
			"reference":   reference,
		})
	})

	// 5. API: Status Polling
	r.GET("/api/Tickets/status/:reference", func(c *gin.Context) {
		ref := c.Param("reference")
		status, exists := ticketTracker[ref]
		if !exists {
			status = "NotFound"
		}
		c.JSON(200, gin.H{"status": status})
	})

	// 6. Webhook Listener
	r.POST("/webhooks/payments", func(c *gin.Context) {
		var payload models.WebhookPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.String(400, "Bad Request")
			return
		}

		// Handle TitleCase 'ReferenceNumber' from OneKhusa Sandbox
		myRef := payload.MetaData.ReferenceNumber
		if myRef == "" {
			myRef = payload.SourceReferenceNumber
		}

		log.Printf("Webhook received for Ref: %s", myRef)

		if payload.ResponseCode == "S100" || payload.TransactionStatusCode == "S" {
			ticketTracker[myRef] = "Paid"
			// Emit real-time update to all dashboard clients
			server.BroadcastToNamespace("/", "webhook_received", gin.H{
				"reference": myRef,
				"status":    "Paid",
			})
		}

		c.String(200, "acknowledged")
	})

	// 7. Socket.io Handlers
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	r.Run(":" + port)
}
