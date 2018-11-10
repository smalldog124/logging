package api_test

import (
	"io/ioutil"
	"logging/route"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListUserHandler_Should_Be_Array_User(t *testing.T) {
	expectUser := `[{"name":"Smalldog","age":"20"},{"name":"Jone","age":"18"}]`
	request := httptest.NewRequest("GET", "/api/v1/user", nil)
	writer := httptest.NewRecorder()
	// api := api.UserAPI{}

	route := route.NewRoute()
	route.ServeHTTP(writer, request)
	response := writer.Result()

	actualUser, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expectUser, string(actualUser))
}
