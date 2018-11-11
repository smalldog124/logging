package api_test

import (
	"github.com/sirupsen/logrus"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"logging/api"
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
	logging := logrus.New()
	logging.SetOutput(buffer)
	api := api.UserAPI{
		Logger: logging,
	}

	route := route.NewRoute(logging, api)
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
	logging := logrus.New()
	logging.SetOutput(buffer)
	api := api.UserAPI{
		Logger: logging,
	}

	route := route.NewRoute(logging, api)
	route.ServeHTTP(writer, request)
	response := writer.Result()

	actualUser, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expectUser, string(actualUser))
}

func Test_CreateUserHandler_Input_UserName_Kaven_And_Age_23_Should_Be_User(t *testing.T) {
	expectUser := `{"id":4,"name":"Kaven","age":"23"}`
	user := api.User{Name: "Kaven", Age: "23"}
	data, _ := json.Marshal(user)
	request := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(data))
	writer := httptest.NewRecorder()
	buffer := &bytes.Buffer{}
	logging := logrus.New()
	logging.SetOutput(buffer)
	api := api.UserAPI{
		Logger: logging,
	}

	route := route.NewRoute(logging, api)
	route.ServeHTTP(writer, request)
	response := writer.Result()

	actualUser, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expectUser, string(actualUser))
}
