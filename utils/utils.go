package utils

import (
	"authentication-service/models/configModel"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Env configModel.ConfigModel

func InitConfig() configModel.ConfigModel {
	err := godotenv.Load()
	if err != nil {
		logrus.Infof("erro ao obter .env: %v", err)
	}

	Env = configModel.ConfigModel{
		DBConf: &configModel.DBConfigModel{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
		},
		EMailConf: &configModel.EMailConfigModel{
			User: os.Getenv("EMAIL_USER"),
			Pass: os.Getenv("EMAIL_PASS"),
			SMTP: os.Getenv("EMAIL_SMTP"),
			Port: os.Getenv("EMAIL_PORT"),
		},
		URLSite: os.Getenv("URL_SITE"),
	}

	return Env
}
