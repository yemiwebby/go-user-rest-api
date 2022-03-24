package configs

import (
	"os"
)

func EnvMongoURI() string {
    return os.Getenv("MONGOURI")
}