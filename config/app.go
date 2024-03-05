package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

var CfgValues Config
var LogFile *log.Logger

type Config struct {
	Env            string `env:"ENV"`
	DBType         string `env:"DB_TYPE" env-default:"sqlite"`
	DBPort         string `env:"DB_PORT" env-default:"5432"`
	DBHost         string `env:"DB_HOST" env-default:"localhost"`
	DBName         string `env:"DB_NAME" env-default:"database"`
	DBUser         string `env:"DB_USER" env-default:"user"`
	DBPassword     string `env:"DB_PASSWORD"`
	RandomSeconds  int    `env:"RANDOM_SECONDS"`
	TelegramToken  string `env:"BOT_TOKEN"`
	TelegramChatID string `env:"CHAT_ID"`
}

func init() {
	LogFile = initLogFile()
	CfgValues = readConfig()
}

func readConfig() Config {
	// Load ENV config file
	var cfg Config

	error := cleanenv.ReadConfig("./.env", &cfg)

	if error != nil {
		errMsg := "Error loading .env file"
		panic(errMsg)
	}

	return cfg
}

func initLogFile() *log.Logger {

	logFile, err := os.OpenFile("./error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return log.New(logFile, "[error] ", log.LstdFlags)
}
