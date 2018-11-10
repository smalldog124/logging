package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func responseOK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, data)
	context.Set("responseBody", data)
}

func ListUserHandler(context *gin.Context) {
	user := []User{
		{
			ID:   1,
			Name: "Smalldog",
			Age:  "20",
		},
		{
			ID:   2,
			Name: "Jone",
			Age:  "18",
		},
	}
	responseOK(context, user)
}

func GetUserHandler(context *gin.Context) {
	user := User{
		ID:   1,
		Name: "Smalldog",
		Age:  "20",
	}
	responseOK(context, user)
}

func CreateUserHandler(context *gin.Context) {
	var user User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		log.Println("can not bind json", err)
		return
	}
	user.ID = 4
	responseOK(context, user)
}
