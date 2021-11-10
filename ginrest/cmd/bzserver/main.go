package main

import (
	"fmt"

	routes "example.com/m/v2/ginrest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("Merhaba yalan dunya")

	routes.AddUserRoutes(router)

	router.Run("localhost:8080")
}
