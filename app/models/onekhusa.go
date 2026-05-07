package models

// Authentication block
type Authentication struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

// Merchant block
type Merchant struct {
	OrganisationID        string `json:"organisationId"`
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
}

// Payment block
type Payment struct {
	SourceReferenceNumber string  `json:"sourceReferenceNumber"`
	Description           string  `json:"description"`
	Amount                float64 `json:"amount"`
}

// Route block
type Route struct {
	SuccessRedirectionUrl string `json:"successRedirectionUrl"`
	FailureRedirectionUrl string `json:"failureRedirectionUrl"`
	CallbackApiUrl        string `json:"callbackApiUrl"`
}

// Full Checkout Request
type CheckoutRequest struct {
	Authentication Authentication `json:"authentication"`
	Merchant       Merchant       `json:"merchant"`
	Payment        Payment        `json:"payment"`
	Route          Route          `json:"route"`
}

// OneKhusa API Response
type CheckoutResponse struct {
	SourceReferenceNumber string `json:"sourceReferenceNumber"`
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
	PaymentTransactionID  string `json:"paymentTransactionId"`
}

// Webhook Payload
type WebhookPayload struct {
	ResponseCode           string `json:"responseCode"`
	TransactionStatusCode  string `json:"transactionStatusCode"`
	SourceReferenceNumber  string `json:"sourceReferenceNumber"`
	MetaData               struct {
		ReferenceNumber string `json:"ReferenceNumber"` // TitleCase from Sandbox
	} `json:"metaData"`
}