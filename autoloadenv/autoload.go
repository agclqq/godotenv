package autoloadenv

import (
	"github.com/agclqq/godotenv"
	"os"
)

func init() {
	env := os.Getenv("GO_ENV")
	fileName := ".env"
	if "" != env {
		fileName += "." + env
	}
	godotenv.Read(fileName)
}
