package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createSystemMetricPayload struct {
	Name string `json:"metric_name"`
}

func (controller *MetricsController) AddSystemMetric(c *gin.Context) {
	var payload createSystemMetricPayload
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

	redisErr := controller.redisClient.AddToList(string(payloadJson), "system-metrics")
	if redisErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": marshalErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "system metric successfully queued"})
}
