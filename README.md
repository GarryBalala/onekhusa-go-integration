# OneKhusa Request to Pay Go (Golang) Integration Reference

A high-performance reference implementation of the **OneKhusa Payment Gateway** using **Go (Golang)** and the **Gin Web Framework**. This project demonstrates the **Hosted Checkout** flow with real-time webhook verification and Socket.io integration for seamless payment processing.

## 📋 Table of Contents

- [Features](#-key-features)
- [Prerequisites](#prerequisites)
- [Installation Guide](#installation-guide)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Webhook Setup](#webhook-setup)
- [API Integration](#api-integration)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## 🚀 Key Features

- **Gin Framework**: Fast and efficient HTTP routing for production-grade performance
- **Strictly Typed Models**: Go structs with JSON tags for precise OneKhusa API mapping
- **Hosted Checkout**: Professional redirection flow to OneKhusa's secure payment page
- **Real-Time Webhooks**: Automated transaction verification using WebSockets
- **Concurrency**: Leverages Go routines for efficient Socket.io handling
- **Environment Configuration**: Secure credential management via .env files
- **Dashboard UI**: Built-in HTML dashboard for payment monitoring

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.19 or higher ([Download Go](https://golang.org/dl/))
- **Git**: For cloning the repository
- **ngrok**: For exposing your local server (for webhook testing) - [Download ngrok](https://ngrok.com/download)
- **OneKhusa Account**: Access to OneKhusa Dashboard with API credentials
  - API Key
  - API Secret
  - Organization ID
  - Merchant Number

## Installation Guide

### Step 1: Clone the Repository

```bash
git clone https://github.com/GarryBalala/onekhusa-go-integration.git
cd onekhusa-go-integration
```

### Step 2: Install Go Dependencies

Download all required Go modules:

```bash
go mod download
```

Or use the alternative command:

```bash
go get ./...
```

### Step 3: Configure Environment Variables

Create a `.env` file in the root directory:

```bash
touch .env
```

Add your configuration variables to the `.env` file:

```env
# OneKhusa API Credentials
ONEKHUSA_API_KEY=your_api_key_here
ONEKHUSA_API_SECRET=your_api_secret_here
ONEKHUSA_ORG_ID=your_org_id_here
ONEKHUSA_MERCHANT_NUMBER=79619974

# OneKhusa API Endpoints
ONEKHUSA_BASE_URL=https://api.onekhusa.com/sandbox/v1
ONEKHUSA_CHECKOUT_URL=https://api.onekhusa.com/sandbox/v1/checkout/rtp/initiate

# Callback Configuration
PUBLIC_CALLBACK_URL=https://your-id.ngrok-free.dev

# Server Port
PORT=8080
```

**Important**: Replace placeholders with your actual OneKhusa credentials from the dashboard.

### Step 4: Verify Installation

Run the following command to verify all dependencies are properly installed:

```bash
go mod verify
```

## 📂 Project Structure

```text
onekhusa-go-integration/
├── app/
│   ├── models/
│   │   └── onekhusa.go                 # Data structures (Structs)
│   └── services/
│       └── onekhusa_service.go         # OneKhusa API Logic
├── public/
│   └── index.html                      # Dashboard (Frontend UI)
├── .env                                # Secrets & Configuration
├── .env.example                        # Example environment file
├── .gitignore                          # Git ignore rules
├── go.mod                              # Go Module definition
├── go.sum                              # Dependency checksums
├── main.go                             # Entry point & Socket setup
└── README.md                           # This file
```

## Configuration

### Environment Variables Explained

| Variable | Description | Example |
|----------|-------------|---------|
| `ONEKHUSA_API_KEY` | Your OneKhusa API key | `pk_live_xxxxx` |
| `ONEKHUSA_API_SECRET` | Your OneKhusa API secret | `sk_live_xxxxx` |
| `ONEKHUSA_ORG_ID` | Your organization ID in OneKhusa | `org_123456` |
| `ONEKHUSA_MERCHANT_NUMBER` | Your merchant number | `79619974` |
| `ONEKHUSA_BASE_URL` | Base URL for API calls | `https://api.onekhusa.com/sandbox/v1` |
| `ONEKHUSA_CHECKOUT_URL` | Checkout initiation endpoint | `https://api.onekhusa.com/sandbox/v1/checkout/rtp/initiate` |
| `PUBLIC_CALLBACK_URL` | Your public callback URL (ngrok or domain) | `https://abc123.ngrok-free.dev` |
| `PORT` | Server port | `8080` |

### Switching Environments

For **Production**:
```env
ONEKHUSA_BASE_URL=https://api.onekhusa.com/v1
ONEKHUSA_CHECKOUT_URL=https://api.onekhusa.com/v1/checkout/rtp/initiate
```

For **Sandbox** (default):
```env
ONEKHUSA_BASE_URL=https://api.onekhusa.com/sandbox/v1
ONEKHUSA_CHECKOUT_URL=https://api.onekhusa.com/sandbox/v1/checkout/rtp/initiate
```

## Running the Application

### Development Mode

Start the server with auto-reload (requires `air` package):

```bash
go install github.com/cosmtrek/air@latest
air
```

### Standard Execution

Start the server:

```bash
go run main.go
```

**Expected Output**:
```
[GIN-debug] Loaded HTML rendering engine
[GIN-debug] Listening and serving HTTP on :8080
```

### Access the Dashboard

Open your browser and navigate to:

```
http://localhost:8080
```

## Webhook Setup

### Setup with ngrok

Webhooks require a public URL to receive payment notifications. Use ngrok to expose your local server:

#### Step 1: Start ngrok

```bash
ngrok http 8080
```

ngrok will generate a public URL like: `https://abc123.ngrok-free.dev`

#### Step 2: Update Configuration

Update the `PUBLIC_CALLBACK_URL` in your `.env` file:

```env
PUBLIC_CALLBACK_URL=https://abc123.ngrok-free.dev
```

#### Step 3: Restart the Application

```bash
go run main.go
```

#### Step 4: Register Webhook in OneKhusa Dashboard

1. Log in to [OneKhusa Dashboard](https://dashboard.onekhusa.com)
2. Navigate to **Settings → Webhooks**
3. Add a new webhook with the endpoint:
   ```
   https://abc123.ngrok-free.dev/webhooks/payments
   ```
4. Select events: `payment.completed`, `payment.failed`, `payment.pending`
5. Save and test the webhook

### Webhook Events

The application handles the following webhook events:

| Event | Description |
|-------|-------------|
| `payment.completed` | Payment successfully processed |
| `payment.failed` | Payment transaction failed |
| `payment.pending` | Payment is pending verification |

## API Integration

### Initiating a Payment

**Endpoint**: `POST /api/payments/initiate`

**Request**:
```json
{
  "amount": 1000,
  "phone_number": "254712345678",
  "reference": "ORD-2026-001",
  "description": "Payment for services"
}
```

**Response**:
```json
{
  "checkout_url": "https://checkout.onekhusa.com/...",
  "session_id": "session_abc123"
}
```

### Checking Payment Status

**Endpoint**: `GET /api/payments/:reference`

**Response**:
```json
{
  "reference": "ORD-2026-001",
  "status": "completed",
  "amount": 1000,
  "timestamp": "2026-05-15T10:30:00Z"
}
```

## Troubleshooting

### Common Issues and Solutions

#### 1. "cannot find package" Error

**Problem**: Dependencies not installed properly

**Solution**:
```bash
go mod tidy
go mod download
```

#### 2. Port Already in Use

**Problem**: Port 8080 is already in use

**Solution**: Change the port in `.env`:
```env
PORT=8081
```

#### 3. Webhook Not Receiving Events

**Problem**: Payment notifications not arriving

**Solution**:
- Verify ngrok tunnel is running: `ngrok http 8080`
- Check `PUBLIC_CALLBACK_URL` matches ngrok URL
- Verify webhook is registered in OneKhusa Dashboard
- Check firewall/network settings

#### 4. "invalid API Key" Error

**Problem**: API credentials are incorrect

**Solution**:
- Verify credentials in `.env` file
- Check credentials in OneKhusa Dashboard
- Ensure you're using the correct environment (sandbox/production)

#### 5. CORS Errors

**Problem**: Cross-origin requests failing

**Solution**: Update CORS configuration in `main.go`:
```go
config := cors.DefaultConfig()
config.AllowOrigins = []string{"http://localhost:3000"}
```

#### 6. "address already in use" Error

**Problem**: Another process is using the configured port

**Solution**:
```bash
# Find process using the port
lsof -i :8080

# Kill the process
kill -9 <PID>

# Or use a different port
PORT=8081 go run main.go
```


## API Reference

For detailed OneKhusa API documentation, visit:
- [OneKhusa Developer Docs](https://docs.onekhusa.com)
- [API Reference](https://docs.onekhusa.com/api)


## Support

For issues or questions:

- Create an [Issue](https://github.com/GarryBalala/onekhusa-go-integration/issues)
- Contact: [support@onekhusa.com](mailto:support@onekhusa.com)
- Documentation: [OneKhusa Docs](https://docs.onekhusa.com)

---

**Last Updated**: May 15, 2026  
**Author**: Garry Balala  
**Repository**: [GarryBalala/onekhusa-go-integration](https://github.com/GarryBalala/onekhusa-go-integration)
