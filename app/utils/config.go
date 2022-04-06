package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Sample string `yaml:"sample`
}

// 설정로드
func LoadConfig() *Config {
	viper.AddConfigPath("utils/profiles")
	viper.SetConfigType("yaml")
	viper.SetConfigName(getEnv())
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("[-] Load configfile error: %v", err)
		panic(err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("[-] Load config Unmarshal error: %v", err)
		panic(err)
	}

	return conf
}

// 환경변수 조회
func getEnv() string {
	mode := os.Getenv("MODE")

	if mode == "" {
		mode = "local"
	}

	return mode
}
