package models

import (
	"acrcloud-scan-tool-golang/acrcloud"
	"acrcloud-scan-tool-golang/checker"
	"acrcloud-scan-tool-golang/logger"
	"github.com/spf13/viper"
)

const componentName = "models.config"

var (
	RecHandler *acrcloud.Recognizer
	Config     *Configuration
)

type Configuration struct {
	Acrcloud struct {
		Host            string `mapstructure:"host"`
		AccessKey       string `mapstructure:"access_key"`
		AccessSecret    string `mapstructure:"access_secret"`
		RecognizeType   string `mapstructure:"recognize_type"`
		StepInSeconds   int    `mapstructure:"step_in_seconds"`
		RecognizeLength int    `mapstructure:"recognize_length"`
	} `mapstructure:"acrcloud"`
	Custom struct {
		IsCustomFile bool `mapstructure:"is_custom_file"`
		Report       struct {
			ReportPath        string   `mapstructure:"report_path"`
			ReportDelimiter   string   `mapstructure:"report_delimiter"`
			FileFields        []string `mapstructure:"file_fields"`
			MusicFields       []string `mapstructure:"music_fields"`
			CustomFilesFields []string `mapstructure:"custom_files_fields"`
			Filters           struct {
				CombineFilter bool `mapstructure:"combine_filter"`
			} `mapstructure:"filters"`
		} `mapstructure:"report"`
	} `mapstructure:"custom"`
}

func init() {
	checker.CheckFilesExists() // check required files

	//set the config name
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	// set default config
	viper.SetDefault("acrcloud.step_in_seconds", 10)
	viper.SetDefault("acrcloud.recognize_type", "audio")
	viper.SetDefault("acrcloud.recognize_length", 10)
	viper.SetDefault("custom.is_custom_file", false)

	// read config from config file
	if err := viper.ReadInConfig(); err != nil {
		logger.LogPanic(componentName, "Error reading config file: ", err)
	}


	err := viper.Unmarshal(&Config)
	if err != nil {
		logger.LogPanic(componentName, "Error to parse the config: ", err)
	}


	configs := map[string]string{
		"access_key":     Config.Acrcloud.AccessKey,
		"access_secret":  Config.Acrcloud.AccessSecret,
		"host":           Config.Acrcloud.Host,
		"recognize_type": Config.Acrcloud.RecognizeType,
	}

	logger.LogWarn(componentName, "Current Config ", configs)
	// set a global rec handler
	RecHandler = acrcloud.NewRecognizer(configs)

}
