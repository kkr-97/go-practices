package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// )

// // // RedisServer struct to manage Redis connection
// type RedisServer struct {
// 	cache *redis.Client
// }

// // NewRedisServer creates a new RedisServer instance
// func NewRedisServer() *RedisServer {
// 	return &RedisServer{}
// }

// // InitCache initializes the Redis connection
// func (rs *RedisServer) InitCache() {
// 	// Configure the Redis connection
// 	rs.cache = redis.NewClient(&redis.Options{
// 		Addr:         "localhost:6379", // Redis server address
// 		DialTimeout:  5 * time.Second,  // Dial timeout for establishing the connection
// 		ReadTimeout:  3 * time.Second,  // Read timeout for Redis requests
// 		WriteTimeout: 3 * time.Second,  // Write timeout for Redis requests
// 	})

// 	// Try to establish the connection
// 	err := rs.cache.Ping(context.Background()).Err()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to Redis: %v", err)
// 	}
// 	fmt.Println("Connected to Redis successfully")

// 	// Start a background goroutine to keep checking the Redis connection
// 	go rs.monitorConnection()
// }

// // monitorConnection continuously monitors the Redis connection status
// func (rs *RedisServer) monitorConnection() {
// 	for {
// 		// Ping the Redis server to check the connection status
// 		err := rs.cache.Ping(context.Background()).Err()
// 		if err != nil {
// 			// If connection failed, log the error and try to reconnect
// 			log.Printf("Redis connection lost: %v. Attempting to reconnect...", err)
// 			// rs.reconnect()
// 		} else {
// 			// Connection is successful, no need to reconnect
// 			log.Println("Redis connection is healthy")
// 		}

// 		// Wait for a certain time before checking again
// 		time.Sleep(10 * time.Second)
// 	}
// }

// // // reconnect attempts to re-establish the Redis connection
// func (rs *RedisServer) reconnect() {
// 	for {
// 		// Attempt to reconnect to Redis
// 		err := rs.cache.Ping(context.Background()).Err()
// 		if err == nil {
// 			// Connection is re-established
// 			log.Println("Redis reconnected successfully")
// 			return
// 		}
// 		// Retry after some delay
// 		log.Printf("Reconnection attempt failed: %v. Retrying in 5 seconds...", err)
// 		time.Sleep(5 * time.Second)
// 	}
// }

// func main() {
// 	// Create a new RedisServer instance
// 	redisServer := NewRedisServer()

// 	// Initialize the Redis connection
// 	redisServer.InitCache()

// 	// Simulate the server running, handling other requests
// 	// The monitoring of the Redis connection will continue in the background
// 	select {}
// }
