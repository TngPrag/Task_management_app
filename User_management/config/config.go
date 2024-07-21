package config

import (
        "log"
        "os"

        "github.com/joho/godotenv"
)

func ReadConfigEnvVariable(key string) string {
        // load .env file
        err := godotenv.Load("config/.env")
        if err != nil {
                log.Fatalf("Error loading .env file")
        }
        return os.Getenv(key)
}
