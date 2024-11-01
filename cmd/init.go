package cmd

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"xiaozhu/router"
	"xiaozhu/utils"
)

const defaultPort = "80"

func Init() {
	if err := InitConf(); err != nil {
		log.Fatalln("配置初始失败：", err)
	}

	// 初始化日志
	if err := InitLogs(); err != nil {
		log.Fatalln("日志初始失败：", err)
	}

	if err := utils.InitMysql(); err != nil {
		log.Fatalln("MYSQL初始失败：", err)
	}

	if err := utils.InitRedis(); err != nil {
		log.Fatalln("redis初始失败：", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件: %s 发生变化,Op %d: \n", e.Name, e.Op)
	})
}

func ServerRun() {
	Init()
	r := router.InitRouter()
	port := viper.GetString("port")
	if port == "" {
		port = defaultPort
	}
	if err := r.Run(":" + port); r != nil {
		log.Fatal("服务启动失败： %w", err)
	}
}
