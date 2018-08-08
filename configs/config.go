package configs

import (
	"github.com/spf13/viper"
	"strings"
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
)

type Config struct {
	Name string
}

func (c *Config)initConfig() error  {

	if c.Name != "" {
		viper.SetConfigName(c.Name)
	}else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("MEETER")

	replacer := strings.NewReplacer(".","_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadConfig();err != nil {
		return err
	}

	return nil
}

func (c *Config)watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed: %s", in.Name)
	})
	
}

func Init(cfg string) error  {

	c := Config{
		Name: cfg,
	}

	if err := c.initConfig();err != nil {
		return err
	}

	c.watchConfig()

	return nil
}
