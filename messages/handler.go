package messages

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveMessage stores a chat message in MongoDB
func SaveMessage(collection *mongo.Collection, msg Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg.Time = time.Now()

	_, err := collection.InsertOne(ctx, msg)
	if err != nil {
		log.Println("Error inserting message into MongoDB:", err)
		return err
	}

	return nil
}
