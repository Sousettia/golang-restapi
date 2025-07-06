package tasks

import (
	"context"
	"log"
	"time"

	"github.com/sousettia/go-restapi/config"
	"go.mongodb.org/mongo-driver/bson"
)

func StartUserCountLogger() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		countUsers()
	}
}

func countUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := config.DB.Collection("users").CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error counting users: %v", err)
		return
	}

	log.Printf("📊 Total users in DB: %d\n", count)
}
