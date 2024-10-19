package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config
var Viper *viper.Viper

type Config struct {
	AesKey      string
	AppID       string
	Token       string
	Wkhtmltopdf string
	Ffmpeg      string
}

func init() {
	Viper, _ = InitConfig()
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")
	if err := v.ReadInConfig(); err == nil {
		// log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}
	return v, nil
}
