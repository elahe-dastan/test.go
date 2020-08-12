package config

import (
	"bytes"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Redis Redis `mapstructure:"redis"`
}

type Redis struct {
	Addr string `mapstructure:"addr"`
}

func Read(filename string) Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	if filename != "" {
		viper.SetConfigFile(filename)
		if err := viper.MergeInConfig(); err != nil {
			log.Warnf("loading config file [%s] failed: %s", filename, err)
		} else {
			log.Infof("config file [%s] loaded successfully", filename)
		}
	}

	viper.SetEnvPrefix("test")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}