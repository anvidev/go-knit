package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func MustLoad() {
  if err := godotenv.Load(); err != nil {
    panic(err)
  }
}

func MustEnv(str string) string {
  env := os.Getenv(str)
  if len(env) == 0 {
    panic(fmt.Sprintf("%s env variable is missing", str))
  }
  return env
}

func IsDevelopment() bool {
  return MustEnv("APP_ENV") == "development"
}

func IsProduction() bool {
  return MustEnv("APP_ENV") == "production"
}
