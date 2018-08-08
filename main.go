package main

import (
	"Meeter/router"
	orm "Meeter/database"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	router := router.Init()
	defer orm.DB.Close()
	router.Run(":8085")
}
