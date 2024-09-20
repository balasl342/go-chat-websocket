package main

import (
	"log"
	"net/http"

	"chat-app/messages"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var chatCollection *mongo.Collection

func main() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Select the database and collection
	chatCollection = client.Database("chat_db").Collection("messages")

	http.HandleFunc("/ws", handleConnections)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		var msg messages.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		// Save message to MongoDB
		err = messages.SaveMessage(chatCollection, msg)
		if err != nil {
			log.Println("Error saving message:", err)
			break
		}

		// Send confirmation back to the client
		err = conn.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing JSON:", err)
			break
		}
	}
}
