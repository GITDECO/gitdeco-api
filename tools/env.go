package tools

import (
	"github.com/joho/godotenv"
)

func EnvFileLoad() error {
	envFilePaths := []string{
		"../.env",
	}
	return godotenv.Load(envFilePaths...)
}
