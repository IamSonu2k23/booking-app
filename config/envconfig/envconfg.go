package envconfig

import (
	"fmt"
	"github.com/joho/godotenv"
)

func InitEnvVars() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("OOPs .env not file found", err)
	}
}
