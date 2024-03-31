package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env        string
	Port       string
	ServerHost string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPass     string
	DbName     string
	SecretKey  string
}

func LoadConfig() (*AppConfig, error) {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln("Error when loading .env", err)
		return nil, err
	}
	dbPort, _ := strconv.Atoi(env["DB_PORT"])
	return &AppConfig{
		Env:    env["ENV"],
		Port:   env["PORT"],
		DbHost: env["DB_HOST"],
		DbPort: dbPort,
		DbUser: env["DB_USER"],
		DbPass: env["DB_PASS"],
		DbName: env["DB_NAME"],
	}, nil
}
