package controllers

import (
	"github.com/gin-gonic/gin"
	"metrics-queue/redis_client"
)

type MetricsController struct {
	redisClient *redis_client.RedisClient
}

func NewMetricsController(client *redis_client.RedisClient) *MetricsController {
	return &MetricsController{redisClient: client}
}

func (controller *MetricsController) SetUp(router gin.IRouter) {
	router.POST("/api/metrics/system", controller.AddSystemMetric)
	router.POST("/api/metrics/training", controller.AddTrainingMetric)
	router.POST("/api/metrics/personal", controller.AddPersonalMetric)
}
