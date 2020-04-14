package common

import (
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHConfig 客户端ssh配置文件
type SSHConfig struct {
	// 认证的用户
	Uses string
	//认证的密码
	Pwd string
	//认证的公钥
	PublicKey string
	// 是否公钥认证
	IsPublicKey bool
	//地址
	Address string
	//端口
	Port string
}

// SSHClient 客户端信息

// SSHSession is a ssh session
func SSHSession(config *SSHConfig) (*ssh.Session, error) {
	clientConfig := &ssh.ClientConfig{
		User: config.Uses,
		Auth: []ssh.AuthMethod{
			//密码认证方式
			ssh.Password(config.Pwd),
		},
		Timeout: 30 * time.Second,
	}
	clientConfig.SetDefaults()
	clientConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	client, err := ssh.Dial("tcp", config.Address+":"+config.Port, clientConfig)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}
