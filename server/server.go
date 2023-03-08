package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"os"
	"rankwillServer/common"
	"rankwillServer/middleware"
	_ "rankwillServer/model"
	"rankwillServer/router"
)

func GinRun() {
	InitConfig()
	_ = common.InitRedis()
	_ = common.InitDB()
	r := gin.Default()
	r.Use(middleware.CORS())
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	panic(r.Run(":" + port)) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
