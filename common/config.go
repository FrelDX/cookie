package common

import (
	"log"
	"strings"

	"github.com/Unknwon/goconfig"
)

//HostInfo 主机用户相关信息
type HostInfo struct {
	Port        string
	Address     string
	User        string
	Pwd         string
	IsPublicKey bool
	PublicKey   string
}

//HostConfig 设置配置文件
type HostConfig struct {
	Path   string
	config map[string]string
}

// Load 加载配置文件
func (p *HostConfig) Load(area string) {
	cfg, err := goconfig.LoadConfigFile(p.Path)
	if err != nil {
		log.Fatal(err)
	}
	host, err := cfg.GetSection(area)
	if err != nil {
		log.Fatal(err)
	}
	p.config = host
}

// ParseHostList 解析原始信息获取所有的主机列表
func (p *HostConfig) ParseHostList() []string {
	hostList := []string{}
	for k := range p.config {
		hostList = append(hostList, k)
	}
	return hostList
}

// Toparse 根据配置文件中的组解析对应的主机信息
func (p *HostConfig) Toparse(host string) (*HostInfo, error) {
	hostDataTmp := p.config[host]
	tmp := strings.Split(hostDataTmp, " ")
	//表示密码认证
	if len(tmp) == 4 {
		return &HostInfo{
			Address: tmp[0],
			Port:    tmp[1],
			User:    tmp[2],
			Pwd:     tmp[3],
		}, nil
	}
	//秘钥认证
	return nil, nil

}
