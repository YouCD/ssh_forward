package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

var (
	ErrUnknownAuthMethod = errors.New("未知的认证方式")
)

func InitSSHClientConfig(username string, serverAuthMethod map[string]interface{}) (*ssh.ClientConfig, error) {
	//nolint:gosec
	config := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	for s, i := range serverAuthMethod {
		//nolint:gocritic
		switch strings.ToLower(s) {
		case "publickeys":
			methodForPrivateKey := ServerAuthMethodForPrivateKey(i)
			config.Auth = methodForPrivateKey
			/*
				// 创建sshClient
				sshClientConn, err := ssh.Dial("tcp", serverAddrString, config)
				if err != nil {
					log.Fatalf("ssh.Dial failed: %s", err)
					os.Exit(1)
				}
			*/

			// 创建连接
			return config, nil
		}
	}
	return nil, ErrUnknownAuthMethod
}

func ServerAuthMethodForPrivateKey(i interface{}) []ssh.AuthMethod {
	var p PublicKeys
	marshal, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(marshal, &p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := os.ReadFile(p.PrivateKeyPath)
	if err != nil {
		log.Fatalf("SshRsa  err: %s", err)
	}

	signer, err := ssh.ParsePrivateKey(file)
	if err != nil {
		log.Fatalf("SshRsa  err: %s", err)
	}
	return []ssh.AuthMethod{ssh.PublicKeys(signer)}
}
