package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//nolint:nonamedreturns
func InitConfig() (forward Forward) {
	// 1. 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = viper.UnmarshalKey("Forward", &forward)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range forward.Project {
		if f.ServerAuthMethod == nil {
			f.ServerAuthMethod = forward.ServerAuthMethod
		}
		if f.ServerUser == "" {
			f.ServerUser = forward.ServerUser
		}

		if f.ServerAddr == "" {
			f.ServerAddr = forward.ServerAddr
		}
	}

	flag := CheckConfig(forward)
	if !flag {
		fmt.Println("配置文件有误")
		os.Exit(1)
	}
	return forward

	// 3. 创建ssh连接
}

// 2. 校验配置文件
func CheckConfig(forward Forward) bool {
	// 简单的判断一下，有没有重复的本地端口
	localPortList := make([]string, 0)
	m := make(map[string]string)

	for _, f := range forward.Project {
		localPortList = append(localPortList, f.LocalAddr)
		m[f.LocalAddr] = f.LocalAddr
	}

	return len(localPortList) == len(m)
}
