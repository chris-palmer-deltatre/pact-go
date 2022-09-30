package main

import (
	"github.com/chris-palmer-deltatre/pact-go/examples/gin/provider"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login/:id", provider.UserLogin)
	router.POST("/users/:id", provider.IsAuthenticated(), provider.GetUser)
	router.Run("localhost:8080")
}
