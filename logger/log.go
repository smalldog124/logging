package logger

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware(logger *logrus.Logger, UUID func() string) gin.HandlerFunc {
	return func(context *gin.Context) {
		startRequestTime := time.Now()
		requestID := UUID()
		requestLog := map[string]interface{}{
			"requestID": requestID,
		}
		if context.Request.Method == http.MethodPost {
			requestBody, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				logger.Error(err)
			}
			jsonBody := make(map[string]interface{})
			json.Unmarshal(requestBody, &jsonBody)
			requestLog["body"] = jsonBody
			context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		logger.Infof("After Request [%s] %s %v", context.Request.Method, context.Request.URL.String(), requestLog)

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
