package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	BaseConfig BaseConfig
}

type BaseConfig struct {
	System struct {
		MaxGoroutines uint64 `json:"maxGoroutines"`
		Host          string `json:"host"`
		Port          string `json:"port"`
		Key           string `json:"key"`
	} `json:"System"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	
	return &c, nil
}
