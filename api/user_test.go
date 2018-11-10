package api_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"logging/route"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListUserHandler_Should_Be_Array_User(t *testing.T) {
	expectUser := `[{"id":1,"name":"Smalldog","age":"20"},{"id":2,"name":"Jone","age":"18"}]`
	request := httptest.NewRequest("GET", "/api/v1/user", nil)
	writer := httptest.NewRecorder()
	buffer := &bytes.Buffer{}
	logging := log.New(buffer, "INFO:", 0)

	route := route.NewRoute(logging)
	route.ServeHTTP(writer, request)
	response := writer.Result()

	actualUser, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expectUser, string(actualUser))
}

func Test_GetUserHandler_Input_UserID_1_Should_Be_User(t *testing.T) {
	expectUser := `{"id":1,"name":"Smalldog","age":"20"}`
	request := httptest.NewRequest("GET", "/api/v1/user/1", nil)
	writer := httptest.NewRecorder()
	buffer := &bytes.Buffer{}
	logging := log.New(buffer, "INFO:", 0)

	route := route.NewRoute(logging)
	route.ServeHTTP(writer, request)
	response := writer.Result()

	actualUser, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expectUser, string(actualUser))
}
