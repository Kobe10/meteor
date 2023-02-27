package core

import (
	"fmt"
	"github.com/spf13/viper"
	"meteor/utils"
)

func Viper(path ...string) *viper.Viper {
	var mainConfig string
	// 读取核心默认文件
	mainConfig = utils.ConfigFile
	fmt.Printf("正在读取主配置文件 , config的路径为%v\n", utils.ConfigFile)

	v := viper.New()
	v.AddConfigPath("conf")
	//v.AddConfigPath("conf")
	v.SetConfigName(mainConfig)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("当前项目目录位置: ", v, "Fatal error config file: %s \\n", err))
	}

	v.WatchConfig()

	//v.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("config file changed:", e.Name)
	//	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
	//		fmt.Println(err)
	//	}
	//})
	//
	//if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
	//	fmt.Println(err)
	//}
	return v
}
