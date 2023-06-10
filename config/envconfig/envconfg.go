package envconfig

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvVars() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load EnvVars: %s", err)
	}
}
