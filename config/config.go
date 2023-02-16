package config

import "github.com/spf13/viper"




type Config struct {
	DBUser      string `mapstructure:"DB_USER"`
	DBPass      string `mapstructure:"DB_PASS"`
	DBHost string `mapstructure:"DB_HOST"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      string `mapstructure:"DB_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
			return
	}

	err = viper.Unmarshal(&config)
	return
}