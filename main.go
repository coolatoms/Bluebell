package main

import (
	"fmt"
	"studyWeb/Bluebell/controller"
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/dao/redis"
	"studyWeb/Bluebell/logger"
	"studyWeb/Bluebell/pkg/snowflake"
	"studyWeb/Bluebell/router"
	"studyWeb/Bluebell/setting"
)

func main() {
	//viper解析
	if err := setting.Init("./config/config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
	}

	//日志库初始化
	fmt.Println(setting.Conf.Mode)
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	//数据库连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Println(err)
		return
	}
	defer mysql.Close()

	//	redis初始化
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		return
	}
	defer redis.Close()

	//雪花生成ID
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Println("snowflake field :err", err)
		return
	}

	//初始化gin框架内置的校验器的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		return
	}

	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
