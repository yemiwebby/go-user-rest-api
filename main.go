package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
        router := gin.Default()
  
		configs.ConnectDB()

		routes.UserRoute(router)

		port := os.Getenv("PORT")

		if port == "" {
			port = "3000"
		}
  
        // router.Run("localhost:6000")
		router.Run(":" + port) 
}