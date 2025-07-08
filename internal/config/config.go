package config

import (
	"finance/pkg/logger"
	"fmt"
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
	log := logger.New("config", true)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Loading Config failed", map[string]string{
			"error": err.Error(),
		})
		return Config{}, err
	}
	DB_config_path := Config{DB_username: os.Getenv("DB_USER"), DB_password: os.Getenv("DB_PASSWORD"), DB_host: os.Getenv("DB_HOST"), DB_port: os.Getenv("DB_PORT"), DB_name: os.Getenv("DB_NAME")}
	return DB_config_path, nil
}

type ConfigServer struct {
	Port string `yaml:"port"`
}

func LoadConfigServer(configPath string) (*ConfigServer, error) {
	log := logger.New("config", true)
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Reading Config file failed", map[string]string{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("не удалось прочитать файл %s: %w", configPath, err)
	}

	var port ConfigServer
	err = yaml.Unmarshal(data, &port)
	if err != nil {
		log.Fatal("Parsing YAML failed", map[string]string{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("не удалось распарсить YAML: %w", err)
	}

	return &ConfigServer{
		Port: port.Port,
	}, nil
}
