package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/go-redis/redis/v8"
// )

// type RedisServer struct {
// 	cache *redis.Client
// }

// var ctx = context.Background()

// func (rs *RedisServer) initRedisClient() {
// 	rs.cache = redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})
// }

// func (rs *RedisServer) subscribe(channelName string) {
// 	// Subscribe to the channel
// 	sub := rs.cache.Subscribe(ctx, channelName)
// 	msg, err := sub.Receive(ctx)
// 	if err != nil {
// 		log.Fatal("Error while subscription:", err)
// 	} else {
// 		fmt.Println("Subscribed to channel:", msg)
// 	}

// 	// Receive messages
// 	ch := sub.Channel()

// 	for msg := range ch {
// 		fmt.Printf("Received message: %s\n", msg.Payload)

// 		if msg.Payload == "REDIS_DOWN" {
// 			fmt.Println("Action: Redis is down. Notify admin or retry connection.")
// 		} else if msg.Payload == "REDIS_UP" {
// 			fmt.Println("Action: Redis is online.")
// 		}
// 	}
// }

// // func (rs *RedisServer) publishMessage(channelName string, message string) {
// // 	err := rs.cache.Publish(ctx, channelName, message).Err()
// // 	if err != nil {
// // 		log.Fatalf("Failed to publish message: %v", err)
// // 	}
// // 	fmt.Printf("Published message: %s to channel: %s\n", message, channelName)
// // }

// func main() {
// 	rs := RedisServer{}

// 	rs.initRedisClient()

// 	// Start the subscriber in a separate goroutine
// 	go rs.subscribe("redis-status")

// 	// Simulate publishing messages
// 	// rs.publishMessage("redis-status", "REDIS_DOWN")
// 	// rs.publishMessage("redis-status", "REDIS_UP")

// 	select {}
// }
