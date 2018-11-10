package logger_test

import (
	"bytes"
	"log"
	"logging/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func UUID() string {
	return "6b66cff4-e0ad-11e8-9820-f40f2430c31d"
}
func Test_ListUser_Method_GET_Should_Be_1_Line_Of_Log_Info(t *testing.T) {
	expected := `INFO:{"requestID":"6b66cff4-e0ad-11e8-9820-f40f2430c31d","status_code":200,"body":{"age":"20","name":"Smalldog"},"response_time":"0.00 ms"}
`
	buffer := &bytes.Buffer{}
	logging := log.New(buffer, "INFO:", 0)

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
	// t.Errorf("actual\n '%s' \n expect\n '%s' ", buffer.String(), expected)
	assert.Equal(t, expected, buffer.String())
}
