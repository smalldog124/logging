package logger_test

import (
	"bytes"
	"logging/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func UUID() string {
	return "6b66cff4-e0ad-11e8-9820-f40f2430c31d"
}
func Test_ListUser_Method_GET_Should_Be_1_Line_Of_Log_Info_Request(t *testing.T) {
	expected := `{"level":"info","msg":"After Request [GET] /users map[requestID:6b66cff4-e0ad-11e8-9820-f40f2430c31d]"}
{"level":"info","msg":"Send Response map[requestID:6b66cff4-e0ad-11e8-9820-f40f2430c31d statusCode:200 body:map[name:Smalldog age:20] responseTime:0.00 ms]"}
`
	buffer := &bytes.Buffer{}
	logging := logrus.New()
	logging.SetOutput(buffer)
	logging.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})

	route := gin.Default()
	route.Use(logger.LoggingMiddleware(logging, UUID))
	route.GET("users", func(context *gin.Context) {
		response := gin.H{"name": "Smalldog", "age": "20"}
		context.JSON(http.StatusOK, response)
		context.Set("statusCode", http.StatusOK)
		context.Set("responseBody", response)
	})

	request, _ := http.NewRequest("GET", "/users", nil)
	writer := httptest.NewRecorder()
	route.ServeHTTP(writer, request)
	// t.Errorf("actual\n '%s'\n expect\n '%s'", buffer.String(), expected)
	assert.Equal(t, expected, buffer.String())
}
