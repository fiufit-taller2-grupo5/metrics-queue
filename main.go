package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"metrics-queue/api/controllers"
	"metrics-queue/redis_client"
	"os"
)

func main() {
	port := portFromEnv()
	router := gin.Default()
	err := initControllers(router)
	if err != nil {
		fmt.Println("Failed initializing controllers: " + err.Error())
		return
	}

	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error running metrics server: " + err.Error())
		return
	}
}

func initControllers(router gin.IRouter) error {
	redisClient := redis_client.NewRedisClient()
	if redisClient == nil {
		return errors.New("failed creating redis client")
	}

	metricsController := controllers.NewMetricsController(redisClient)
	metricsController.SetUp(router)
	return nil
}

func portFromEnv() int {
	if os.Getenv("environment") == "production" {
		return 80
	} else {
		return 8004
	}
}
