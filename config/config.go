package config

import "github.com/spf13/viper"


type DB struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	Username string `mapstructure:"db_user"`
	Password string `mapstructure:"db_pass"`
	DBName   string `mapstructure:"db_name"`
}

func LoadConfig() (config DB, err error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
