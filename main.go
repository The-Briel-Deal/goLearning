package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	)

type person struct {
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Age       int    `json:"age"`
}

var people = []person{
  {FirstName: "Gabriel", LastName: "Ford", Age: 5},
  {FirstName: "Yaya", LastName: "DaGiwlfwend", Age: 5},
  {FirstName: "Gabriela", LastName: "Ford", Age: 5},
  {FirstName: "Yayoid", LastName: "DaGiwlfwend", Age: 5},
}

func getAlbums(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, people)
}

func main() {
    router := gin.Default()
    fmt.Println("Starting on localhost:8080")
    router.GET("/people", getAlbums)
    router.Run("localhost:8080")
}

