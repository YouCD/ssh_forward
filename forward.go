package main

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func BindConn(localConn net.Conn, serverAddr, remoteAddr string, config *ssh.ClientConfig) {
	sshClientConn, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}

	sshConn, err := sshClientConn.Dial("tcp", remoteAddr)

	// 读取本地数据，发送到远端
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()

	// 读取远端数据，发送到本地
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}
