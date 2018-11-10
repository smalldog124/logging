package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
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
	context.JSON(http.StatusOK, user)
}
