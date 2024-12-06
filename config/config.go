package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl  string `mapstructure:"DBURL"`
	DBname string `mapstructure:"DBNAME"`

	GrpcPort  string `mapstructure:"GRPCPORT"`
	KafkaPort string `mapstructure:"KAFKA_BROKER"`
	KafkaTpic string `mapstructure:"KAFKA_TOPIC"`
	ChatPort string `mapstructure:"CHATPORT"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}