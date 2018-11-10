package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseLog struct {
	RequestID    string      `json:"requestID"`
	StatusCode   int         `json:"status_code"`
	Body         interface{} `json:"body"`
	ResponseTime string      `json:"response_time"`
}

func LoggingMiddleware(logger *log.Logger, UUID func() string) gin.HandlerFunc {
	return func(context *gin.Context) {
		startRequestTime := time.Now()
		requestID := UUID()

		context.Next()

		durationTime := time.Since(startRequestTime)
		responeBody, _ := context.Get("responseBody")
		responseLog := ResponseLog{
			RequestID:    requestID,
			StatusCode:   context.Writer.Status(),
			Body:         responeBody,
			ResponseTime: fomatTimeToMillisecond(durationTime),
		}
		responseLogJson, _ := json.Marshal(responseLog)
		logger.Printf("Send Response %s", string(responseLogJson))
	}
}

func NewUUID() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}

func fomatTimeToMillisecond(timemer time.Duration) string {
	timeFolat64 := float64(timemer / time.Millisecond)
	return fmt.Sprintf("%.2f ms", timeFolat64)
}
