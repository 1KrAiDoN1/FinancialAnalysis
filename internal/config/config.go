package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB_username string //`json:"username"`
	DB_password string //`json:"password"`
	DB_host     string //`json:"host"`
	DB_port     string //`json:"port"`
	DB_name     string // `json:"name"`
}

func SetConfig() (Config, error) {
	err := godotenv.Load("./internal/storages/database/DB_Config.env")
	if err != nil {
		log.Println("Ошибка при чтении конфигурации базы данных")
		return Config{}, err
	}
	DB_config_path := Config{DB_username: os.Getenv("DB_USER"), DB_password: os.Getenv("DB_PASSWORD"), DB_host: os.Getenv("DB_HOST"), DB_port: os.Getenv("DB_PORT"), DB_name: os.Getenv("DB_NAME")}
	return DB_config_path, nil
}

type ConfigServer struct {
	Port string `yaml:"port"`
}

func LoadConfigServer(configPath string) (*ConfigServer, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл %s: %w", configPath, err)
	}

	var port ConfigServer
	err = yaml.Unmarshal(data, &port)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить YAML: %w", err)
	}

	return &ConfigServer{
		Port: port.Port,
	}, nil
}
