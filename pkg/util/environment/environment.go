package environment

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/sholludev/sampoerna_notification/pkg/file"
)

func Init() {
	var err error
	mode := flag.String("mode", "dev", "dev, prod or stage")
	flag.Parse()

	os.Setenv("PROJECT_DIR", "sampoerna_notification") // should change this first for new projects
	rootPath := file.GetRootDirectory()

	switch *mode {
	case "dev":
		err = godotenv.Load(rootPath + "/.env.development")
	case "prod":
		err = godotenv.Load(rootPath + "/.env.production")
	case "stage":
		err = godotenv.Load(rootPath + "/.env.staging")
	default:
		err = godotenv.Load(rootPath + "/.env.development")
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
