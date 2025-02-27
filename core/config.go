package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const configPath = "./config"
const configName = "config"

type config struct {
	Config struct {
		CheckValue int `mapstructure:"checkvalue" validate:"required"`
	} `mapstructure:"config"`
	Mail struct {
		SendMail     bool `mapstructure:"sendmail" validate:"required"`
		MaxSendEvery int  `mapstructure:"maxsendevery" validate:"required"`
		Smtp         struct {
			Host     string `mapstructure:"host" validate:"required"`
			Port     int    `mapstructure:"port" validate:"required"`
			Username string `mapstructure:"username" validate:"required"`
			Password string `mapstructure:"password" validate:"required"`
		} `mapstructure:"smtp"`
		Recipients []string `mapstructure:"recipients" validate:"required"`
		Sender     string   `mapstructure:"sender" validate:"required"`
	} `mapstructure:"mail"`
}

var AppConfig *config

func SetupConfig() {
	// Set up Viper configuration
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Error reading config file: ", err)
	}

	var c config

	if err := viper.Unmarshal(&c); err != nil {
		logrus.Fatal("Error unmarshaling config: ", err)
	}

	// Validate the config
	validateConfig(c)

	AppConfig = &c
}

func validateConfig(c config) {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		logrus.Fatal("Error validating config: ", err)
	}

	logrus.Info("Configuration validated successfully")
}
