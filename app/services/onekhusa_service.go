package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"onekhusa-go/app/models"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type OneKhusaService struct{}

func (s *OneKhusaService) InitiateHostedCheckout(amount float64, reference string) (*models.CheckoutResponse, error) {
	merchantNo, _ := strconv.Atoi(os.Getenv("ONEKHUSA_MERCHANT_NUMBER"))
	callbackBase := os.Getenv("PUBLIC_CALLBACK_URL")

	// 1. Build the nested payload
	payload := models.CheckoutRequest{
		Authentication: models.Authentication{
			APIKey:    os.Getenv("ONEKHUSA_API_KEY"),
			APISecret: os.Getenv("ONEKHUSA_API_SECRET"),
		},
		Merchant: models.Merchant{
			OrganisationID:        os.Getenv("ONEKHUSA_ORG_ID"),
			MerchantAccountNumber: merchantNo,
		},
		Payment: models.Payment{
			SourceReferenceNumber: reference,
			Description:           "OneTicket Go Showcase",
			Amount:                amount,
		},
		Route: models.Route{
			SuccessRedirectionUrl: fmt.Sprintf("%s/?ref=%s", callbackBase, reference),
			FailureRedirectionUrl: fmt.Sprintf("%s/?ref=%s", callbackBase, reference),
			CallbackApiUrl:        fmt.Sprintf("%s/webhooks/payments", callbackBase),
		},
	}

	// 2. Convert to JSON
	jsonPayload, _ := json.Marshal(payload)

	// 3. Prepare Request
	url := os.Getenv("ONEKHUSA_CHECKOUT_URL")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// 4. Set Headers
	idempotencyKey := fmt.Sprintf("GO-CHK-%s-%s", reference, uuid.New().String()[:8])
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Idempotency-Key", idempotencyKey)

	// 5. Execute
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OneKhusa API Error: %s", resp.Status)
	}

	// 6. Decode Response
	var result models.CheckoutResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}