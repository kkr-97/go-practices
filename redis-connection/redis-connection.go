package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisServer struct {
	cache *redis.Client
}

var options = &redis.Options{
	Addr: "localhost:6379",
}

func configure() *redis.Client {
	return redis.NewClient(options)
}

func (rs *RedisServer) initCache() {
	rs.cache = configure()

	go rs.monitorConnection()
}

func (rs *RedisServer) monitorConnection() {
	ctx := context.Background()
	for {
		err := rs.cache.Ping(ctx).Err()
		if err == nil {
			fmt.Println("Redis connection is active")
		} else {
			fmt.Println("Redis connection is down, trying to reconnect...")

			ok := rs.reconnect()
			if !ok {
				closeApp()
				return
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func closeApp() {
	// Perform any cleanup tasks here
	fmt.Println("Shutting down the application...")
	os.Exit(500)
}

func (rs *RedisServer) reconnect() bool {
	reconnectTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	rs.cache = configure()
	for {
		select {
		case <-reconnectTimeout.Done():
			fmt.Println("Failed to reconnect to Redis")
			return false
		default:
			if err := rs.cache.Ping(context.Background()).Err(); err == nil {
				fmt.Println("Redis connection re-established")
				return true
			}
		}
		fmt.Println("Reconnect attempt failed, retrying...")
		time.Sleep(2 * time.Second)
	}

}

func main() {
	rs := RedisServer{}
	rs.initCache()

	select {}
}
