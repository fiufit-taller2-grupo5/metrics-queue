package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createPersonalMetricPayload struct {
	MetricType string `json:"type"`
	UserId     int    `json:"user_id"`
	MetricName string `json:"metric"`
	Amount     int    `json:"amount"`
}

func (controller *MetricsController) AddPersonalMetric(c *gin.Context) {
	var payload createPersonalMetricPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payloadJson, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": marshalErr.Error()})
		return
	}

	redisErr := controller.redisClient.AddToList(string(payloadJson), "personal-metrics")
	if redisErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": marshalErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "personal metric successfully queued"})
}
