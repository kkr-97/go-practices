package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/example", func(c *gin.Context) {
		ctx, _ := context.WithTimeout(c.Request.Context(), 4*time.Second)

		time.Sleep(time.Second * 4)
		var wg sync.WaitGroup

		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("Operation completed")
				c.JSON(200, gin.H{
					"message": "Request processed",
				})
			case <-ctx.Done():
				fmt.Println("Operation cancelled:", ctx.Err())
				c.JSON(200, gin.H{
					"message": "Request Timeout",
				})
			}
		}(ctx, &wg)
		wg.Wait()

	})

	r.Run(":8080")
}
