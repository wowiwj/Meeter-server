package main

import (
	"Meeter/router"
	orm "Meeter/database"
	"github.com/spf13/pflag"
	"Meeter/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := configs.Init(*cfg);err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	router := router.Init()
	defer orm.DB.Close()
	router.Run(":8085")
}
