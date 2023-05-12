package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"metrics-queue/api/controllers"
	"metrics-queue/redis_client"
)

const Port = 8004

func main() {
	router := gin.Default()
	initControllers(router)
	err := router.Run(fmt.Sprintf(":%d", Port))
	if err != nil {
		fmt.Println("Error running metrics server: " + err.Error())
	}

}

func initControllers(router gin.IRouter) {
	redisClient := redis_client.NewRedisClient()
	metricsController := controllers.NewMetricsController(redisClient)
	metricsController.SetUp(router)
}
