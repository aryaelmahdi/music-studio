package config

import (
	"log"
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

type SMTP struct {
	EmailHost     string
	EmailPort     string
	EmailUsername string
	EmailPassword string
	From          string
}

func InitConfig() (*Config, *SMTP) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	res := new(Config)
	res, smtpRes := loadConfig()
	if res == nil {
		logrus.Fatal("Cannot load configuration")
		return nil, nil
	}
	return res, smtpRes
}

func loadConfig() (*Config, *SMTP) {
	res := new(Config)
	smtpRes := new(SMTP)
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
		jsonBytes := []byte(val)

		file, err := os.Create("credentials.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.Write(jsonBytes)
		if err != nil {
			panic(err)
		}
	}

	if val, found := os.LookupEnv("EMAIL_HOST"); found {
		smtpRes.EmailHost = val
	}
	if val, found := os.LookupEnv("EMAIL_PORT"); found {
		smtpRes.EmailPort = val
	}
	if val, found := os.LookupEnv("EMAIL_USERNAME"); found {
		smtpRes.EmailUsername = val
	}
	if val, found := os.LookupEnv("EMAIL_PASSWORD"); found {
		smtpRes.EmailPassword = val
	}
	if val, found := os.LookupEnv("FROM"); found {
		smtpRes.From = val
	}

	return res, smtpRes
}
