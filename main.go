package main

func main() {
	// 1. 读取配置文件
	// 2. 校验配置文件
	config := InitConfig()

	NewForwardTask(config)
	select {}
}
