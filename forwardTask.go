package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func NewForwardTask(forwards Forward) {
	yellow := color.New(color.FgYellow).PrintfFunc()

	yellow("如果要想其他设备访问还请 将 '/etc/ssh/sshd_config' 配置文件修改为 'GatewayPorts yes' \n")
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for _, f := range forwards.Project {
		fmt.Printf("项目: %s   %s -> %s -> 远端服务(%s)\n", red(f.Project), f.LocalAddr, f.ServerAddr, cyan(f.RemoteAddr))
		go func(f *forward) {
			// 3.  创建ssh登陆配置
			clientConfig, err := InitSSHClientConfig(f.ServerUser, f.ServerAuthMethod)
			if err != nil {
				log.Fatalf("InitSSHClientConfig failed: %s", err)
			}

			// 4. 创建本地端口监听
			listener, err := NewListener(f.LocalAddr)
			if err != nil {
				log.Fatalf("NewListener failed: %s", err)
			}
			for {
				// 将本地端口接收到的请求转发到远程端口
				localConn, err := listener.Accept()
				if err != nil {
					log.Fatalf("listen.Accept failed: %v", err)
				}
				// 5. 端口转发
				go BindConn(localConn, f.ServerAddr, f.RemoteAddr, clientConfig)
			}
		}(f)
	}
}
