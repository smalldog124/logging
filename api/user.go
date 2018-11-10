package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func responseOK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
	context.Set("responseBody", data)
}

func ListUser(context *gin.Context) {
	user := []User{
		{
			Name: "Smalldog",
			Age:  "20",
		},
		{
			Name: "Jone",
			Age:  "18",
		},
	}
	responseOK(context, user)
}
