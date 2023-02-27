package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"meteor/global"
	"meteor/utils"
)

func EnvViper(path ...string) *viper.Viper {
	var envConfig string
	// 读取环境配置文件

	v := viper.New()
	// 根据配置文件环境 读取对应环境配置文件
	env := global.GVA_VP.GetString("run_mode")
	envConfig = utils.GetEnvConfig(env)
	fmt.Printf("正在读取主配置文件 , config的路径为%v\n", utils.EnvConfig)
	v.AddConfigPath("conf")
	//v.AddConfigPath("conf")
	v.SetConfigName(envConfig)
	v.SetConfigType("yaml")
	err1 := v.ReadInConfig()
	if err1 != nil {
		panic(fmt.Errorf("当前项目目录位置: ", v, "Fatal error config file: %s \\n", err1))
	}

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
