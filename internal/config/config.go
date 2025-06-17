package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_username string //`json:"username"`
	DB_password string //`json:"password"`
	DB_host     string //`json:"host"`
	DB_port     string //`json:"port"`
	DB_name     string // `json:"name"`
}

func SetConfig() (Config, error) {
	err := godotenv.Load("/Users/pavelvasilev/Desktop/FinancialAnalysis/internal/storages/database/DB_Config.env")
	if err != nil {
		log.Println("Ошибка при чтении конфигурации базы данных")
		return Config{}, err
	}
	DB_config_path := Config{DB_username: os.Getenv("DB_USER"), DB_password: os.Getenv("DB_PASSWORD"), DB_host: os.Getenv("DB_HOST"), DB_port: os.Getenv("DB_PORT"), DB_name: os.Getenv("DB_NAME")}
	return DB_config_path, nil
}
