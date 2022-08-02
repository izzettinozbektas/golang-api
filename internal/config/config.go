package config

import "github.com/spf13/viper"

type Config struct {
	DBNAME          string `mapstructure:"DB_NAME"`
	DBUNAME         string `mapstructure:"DB_UNAME"`
	DBPASS          string `mapstructure:"DB_PASS"`
	DBHOST          string `mapstructure:"DB_HOST"`
	REDISURL        string `mapstructure:"REDIS_URL"`
	REDISPASS       string `mapstructure:"REDIS_PASSWORD"`
	GetQuoteFromAPI string `mapstructure:"QUOTE_API_URL"`
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
