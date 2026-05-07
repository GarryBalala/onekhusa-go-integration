# OneKhusa Request to Pay Go (Golang) Integration Reference

A high-performance reference implementation of the **OneKhusa Payment Gateway** using **Go (Golang)** and the **Gin Web Framework**. This project demonstrates the **Hosted Checkout** flow with real-time webhook automation via **go-socket.io**.

## 🚀 Key Features
- **Gin Framework**: Fast and efficient HTTP routing.
- **Strictly Typed Models**: Go structs with JSON tags for precise OneKhusa API mapping.
- **Hosted Checkout**: Professional redirection flow to OneKhusa’s secure payment page.
- **Real-Time Webhooks**: Automated transaction verification using WebSockets.
- **Concurrency**: Leverages Go routines for efficient Socket.io handling.

## 📂 Project Structure
```text
onekhusa-go-integration/
├── app/
│   ├── models/             # Data structures (Structs)
│   │   └── onekhusa.go
│   └── services/           # OneKhusa API Logic
│       └── onekhusa_service.go
├── public/
│   └── index.html          # Dashboard (Frontend)
├── .env                    # Secrets & Configuration
├── .gitignore              # Security
├── go.mod                  # Go Module definition
└── main.go                 # Entry point & Socket setup
🛠️ Setup & Installation
Clone and Enter Directory:
code
Bash
git clone https://github.com/GarryBalala/onekhusa-go-integration.git
cd onekhusa-go-integration
Install Dependencies:
code
Bash
go mod download
Configure Environment (.env):
Create a .env file in the root and add your OneKhusa credentials:
code
Env
ONEKHUSA_API_KEY=your_key
ONEKHUSA_API_SECRET=your_secret
ONEKHUSA_ORG_ID=your_org_id
ONEKHUSA_MERCHANT_NUMBER=79619974
ONEKHUSA_BASE_URL=https://api.onekhusa.com/sandbox/v1
ONEKHUSA_CHECKOUT_URL=https://api.onekhusa.com/sandbox/v1/checkout/rtp/initiate
PUBLIC_CALLBACK_URL=https://your-id.ngrok-free.dev
PORT=8080
Start the Server:
code
Bash
go run main.go
📡 Webhook Setup (NGrok)
Run ngrok http 8080.
Update PUBLIC_CALLBACK_URL in your .env.
Register https://your-id.ngrok-free.dev/webhooks/payments in the OneKhusa Portal.
