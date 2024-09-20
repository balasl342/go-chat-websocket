# Chat Application using WebSocket and MongoDB

This is a simple Go application that demonstrates a chat system using WebSocket for messaging and MongoDB for message persistence.

## Features

- Messaging with WebSocket.
- Messages are stored in MongoDB.
- Clients can send and receive messages in real-time.

## Prerequisites

- Go 1.16 or higher installed.
- MongoDB installed and running locally or accessible via a connection URI.
- Web browser for accessing the WebSocket client (HTML file).

## Installation

1. Clone the repository:

```bash
git clone https://github.com/balasl342/go-chat-websocket.git
```

2. Install Go dependencies:

```bash
go mod tidy
```

3. Ensure MongoDB is running locally on port `27017` or update the MongoDB connection URI in `main.go` if using a cloud instance (e.g., MongoDB Atlas).
