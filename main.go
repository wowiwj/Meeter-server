package main

import (
	"github.com/spf13/pflag"
	"Meeter/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"Meeter/router"
	"github.com/lexkong/log"
	"net/http"
	"fmt"
	"Meeter/models"
)

var (
	cfg = pflag.StringP("config", "c", "", "meeter config file path.")
)

func main() {

	pflag.Parse()

	// 初始化配置
	if err := configs.Init(*cfg);err != nil {
		panic(err)
	}

	// 初始化数据库
	models.DB.Init()
	defer models.DB.Close()

	// 配置gin
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()
	middleware := []gin.HandlerFunc{}

	router.Load(
		// 内核
		g,
		// 中间件
		middleware...,
	)

	addr := viper.GetString("addr")
	fmt.Println(addr)

	// Ping the server to make sure the router is working.
	log.Infof("Start to listening the incoming requests on http address: %s", addr)
	log.Infof(http.ListenAndServe(addr, g).Error())
}

