package initialization

import (
	"wechat-imitation/config"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to read configuration file")
	}
	viper.Unmarshal(&config.Conf)
}
