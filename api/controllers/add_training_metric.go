package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type trainingData struct {
	Calories *int `json:"calories,omitempty"`
	Steps    *int `json:"steps,omitempty"`
	Distance *int `json:"distance,omitempty"`
}

type completeTrainingMetricPayload struct {
	MetricName   string       `json:"metric_name"`
	UserId       int          `json:"user_id"`
	TrainingId   int          `json:"training_id"`
	TrainingData trainingData `json:"data,omitempty"`
}

func (controller *MetricsController) AddTrainingMetric(c *gin.Context) {
	var payload completeTrainingMetricPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(fmt.Sprintf("Received training metric: %+v", payload))

	payloadJson, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": marshalErr.Error()})
		return
	}

	redisErr := controller.redisClient.AddToList(string(payloadJson), "training-metrics")
	if redisErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": marshalErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "training metric successfully queued"})
}
