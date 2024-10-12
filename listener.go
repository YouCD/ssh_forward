package main

import (
	"net"
)

// 搞一个本地端口监听
func NewListener(localAddrString string) (net.Listener, error) {
	// todo 检查端口是否备占用
	//nolint:wrapcheck
	return net.Listen("tcp", localAddrString)
}
