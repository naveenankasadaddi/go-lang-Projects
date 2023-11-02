package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tudo struct {
	Id        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var tudos = []tudo{
	{Id: "1", Item: "Clean room", Completed: false},
	{Id: "2", Item: "Read the book", Completed: true},
	{Id: "3", Item: "meet Naeko", Completed: true},
}

func getTudos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tudos)
}

func main() {
	router := gin.Default()
	router.GET("/tudos", getTudos)
	router.Run("localhost:8000")

	//fmt.Println("Hello ", todos)
}
