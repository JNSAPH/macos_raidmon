package core

import (
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ConfigPath string
var ConfigName string

const DefaultConfigName = "config"
const DefaultConfigPath = "./config"

type config struct {
	Config struct {
		CheckValue int `mapstructure:"checkvalue" validate:"required"`
	} `mapstructure:"config"`
	Mail struct {
		SendMail         bool   `mapstructure:"sendmail" validate:"required"`
		MaxSendEvery     int    `mapstructure:"maxsendevery" validate:"required"`
		DailyReportChron string `mapstructure:"dailyreportchron" validate:"required"`
		Smtp             struct {
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

	// Seperate file name and path into ConfigPath and ConfigName
	ConfigPath, ConfigName = splitPath(ConfigPath)

	logrus.Info("Config path: ", ConfigPath)
	logrus.Info("Config name: ", ConfigName)

	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)

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

func splitPath(fullPath string) (string, string) {
	dir := filepath.Dir(fullPath)
	file := strings.TrimSuffix(filepath.Base(fullPath), filepath.Ext(fullPath))
	return dir, file
}
