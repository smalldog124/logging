package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResponseLog struct {
	RequestID    string      `json:"requestID"`
	StatusCode   int         `json:"status_code"`
	Body         interface{} `json:"body"`
	ResponseTime string      `json:"response_time"`
}

func LoggingMiddleware(logger *logrus.Logger, UUID func() string) gin.HandlerFunc {
	return func(context *gin.Context) {
		startRequestTime := time.Now()
		requestID := UUID()

		context.Next()

		durationTime := time.Since(startRequestTime)
		responeBody, _ := context.Get("responseBody")
		responseLog := map[string]interface{}{
			"requestID":    requestID,
			"statusCode":   context.Writer.Status(),
			"body":         responeBody,
			"responseTime": fomatTimeToMillisecond(durationTime),
		}
		logger.Infof("Send Response %v", responseLog)
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
