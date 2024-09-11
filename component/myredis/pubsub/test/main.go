package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Subscribe to a channel
	pubsub := client.Subscribe(ctx, "mychannel")
	defer pubsub.Close()

	// Check for errors during subscription
	smt, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	} else {
		log.Println(smt)
	}

	// Handle signals to gracefully exit
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Goroutine to handle incoming messages
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Received message: %s\n", msg.Payload)
		}
	}()

	// Publish a message every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Publish a message to the channel
			err := client.Publish(ctx, "mychannel", "Hello, The world !").Err()
			if err != nil {
				panic(err)
			}
		case <-signals:
			fmt.Println("Received interrupt signal. Exiting...")
			return
		}
	}
}
