package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
)

func main() {
        router := gin.Default()
		// router.Use(cors.Default())
		// router.Use(cors.New(cors.Config{
		// 	AllowOrigins: []string{"*"},
		// 	AllowMethods: []string{"PUT", "PATCH", "POST"},
		// 	AllowHeaders: []string{"Origin"},
		// 	ExposeHeaders: []string{"Content-Length"},
		// 	AllowCredentials: true,
		// 	AllowOriginFunc: func (origin string) bool  {
		// 		return origin == ""
		// 	},
		// }))

		corsConfig := cors.DefaultConfig()

		corsConfig.AllowOrigins = []string{"*"}
		// To be able to send tokens to the server.
		corsConfig.AllowCredentials = true

		// OPTIONS method for ReactJS
		corsConfig.AddAllowMethods("OPTIONS")

		// Register the middleware
		router.Use(cors.New(corsConfig))
  
		configs.ConnectDB()

		routes.UserRoute(router)

		port := os.Getenv("PORT")

		if port == "" {
			port = "3000"
		}
  
		router.Run(":" + port) 
}