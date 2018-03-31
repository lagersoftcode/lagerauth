package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Bind     string
	DBConfig struct {
		User     string
		Pass     string
		Database string
		Port     string
		Host     string
	}
	EmailConfig struct {
		Email string
		Pass  string
		Host  string
		Port  string
	}
	JWTKey string
}

func Read() *Config {
	log.Printf("Reading conf from file.\n")

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Error reading configuration:", err)
	}

	return &configuration
}
