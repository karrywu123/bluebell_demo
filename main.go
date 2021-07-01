package main

import (
	"bluebell_demo/dao/mysql"
	"bluebell_demo/dao/redis"
	"bluebell_demo/logger"
	"bluebell_demo/pkg/snowflake"
	"bluebell_demo/routers"
	"bluebell_demo/settings"
	"fmt"
)

func main(){
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	//初始化mysql数据库配置
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	//初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(0); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := routers.SetupRouter()

	//err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	err := r.Run(fmt.Sprintf(":%v", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

