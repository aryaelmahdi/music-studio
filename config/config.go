package config

import (
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SECRET      string
	ProjectID   string
	SDKPath     string
	DatabaseURL string
	MDServerKey string
	MDClientKey string
}

func InitConfig() *Config {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	res := new(Config)
	res = loadConfig()
	if res == nil {
		logrus.Fatal("Cannot load configuration")
		return nil
	}
	return res
}

func loadConfig() *Config {
	res := new(Config)
	if val, found := os.LookupEnv("SECRET"); found {
		res.SECRET = val
	}
	if val, found := os.LookupEnv("ProjectID"); found {
		res.ProjectID = val
	}
	if val, found := os.LookupEnv("SDKPath"); found {
		res.SDKPath = val
	}
	if val, found := os.LookupEnv("DatabaseURL"); found {
		res.DatabaseURL = val
	}
	if val, found := os.LookupEnv("MidtransSandBoxServerKey"); found {
		res.MDServerKey = val
	}
	if val, found := os.LookupEnv("MidtransSandBoxClientKey"); found {
		res.MDClientKey = val
	}
	if val, found := os.LookupEnv("GOOCREDS"); found {

		file, err := os.Create("credentials.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.WriteString(file, val)
		if err != nil {
			panic(err)
		}
	}
	return res
}
