package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web-edu/routers"
	"go-web-edu/service"
	"os"
)

func main() {
	InitConfig()
	// 连接数据库
	db := service.InitDB()
	defer db.Close()

	r := gin.Default()

	// 加载路由
	routers.LoadCategory(r)
	routers.LoadArticle(r)

	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	} else {
		if err := r.Run(); err != nil {
			fmt.Println("startup service failed, err:%v\n", err)
		}
	}

}

// 初始化配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app")               //设置文件名
	viper.SetConfigType("yml")               // 设置文件格式
	viper.AddConfigPath(workDir + "/config") // 文件路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
