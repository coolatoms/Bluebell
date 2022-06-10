package main

import (
	"fmt"
	"studyWeb/Bluebell/controller"
	"studyWeb/Bluebell/pkg/snowflake"
	"studyWeb/Bluebell/router"
	"studyWeb/Bluebell/setting"
)

func main() {
	//viper解析
	if err := setting.Init("./config/config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
	}
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
	r := router.SetupRouter()
	if err := r.Run(); err != nil {
		return
	}
	//err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	//if err != nil {
	//	fmt.Printf("run server failed, err:%v\n", err)
	//	return
	//}
}
