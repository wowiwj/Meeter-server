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

	if err := viper.ReadInConfig();err != nil {
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

func (c *Config)initLog()  {
	passLagerCfg := log.PassLagerCfg {
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

func Init(cfg string) error  {

	c := Config{
		Name: cfg,
	}

	// 初始化日志包
	if err := c.initConfig();err != nil {
		return err
	}

	// 初始化日志文件
	c.initLog()

	c.watchConfig()

	return nil
}
