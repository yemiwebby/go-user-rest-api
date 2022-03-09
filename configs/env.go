package configs

import (
	"os"
)

func EnvMongoURI() string {
    // if os.Getenv("APP_ENV") != "production" {
    //     err := godotenv.Load()
    //     if err != nil {
    //         log.Fatal("Error loading .env file within the application")
    //     }
      
    // }

    return os.Getenv("MONGOURI")
}