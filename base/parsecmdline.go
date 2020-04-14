package base

import (
	"bytes"
	"cookie/common"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

//解析命令行参数
var grep = flag.String("grep", "", "指定一个运行任务的组")
var mod = flag.String("mod", "cmd", "设置一个运行任务的模块，默认cmd")
var cmd = flag.String("cmd", " ", "执行的命令")
var sshConfig *common.SSHConfig

type CmdLine struct {
	Grep string
	Mod  string
	Cmd  string
}

func Parse() *CmdLine {
	flag.Parse()
	if *grep == "" {
		flag.Usage()
		os.Exit(127)
	}
	return &CmdLine{
		Grep: *grep,
		Mod:  *mod,
		Cmd:  *cmd,
	}
}

func Start() {
	para := Parse()

	hostConfig := common.HostConfig{
		Path: "config",
	}
	//加载配置文件的区
	hostConfig.Load(para.Grep)
	//获取怂所以的主机
	wg := sync.WaitGroup{}
	var hostInfo *common.HostInfo
	for _, k := range hostConfig.ParseHostList() {
		hostInfo, _ = hostConfig.Toparse(k)
		wg.Add(1)
		go func(hostInfo *common.HostInfo) {
			sshConfig = &common.SSHConfig{
				Uses:    hostInfo.User,
				Pwd:     hostInfo.Pwd,
				Address: hostInfo.Address,
				Port:    hostInfo.Port,
			}
			ss, err := common.SSHSession(sshConfig)
			if err != nil {
				log.Fatal(err)
			}
			//后台运行，后期需要对输出做统一处理
			Stdout := new(bytes.Buffer)
			ss.Stdout = Stdout
			ss.Stderr = os.Stderr
			//cmdline, err := ss.StdinPipe()
			ss.Run(para.Cmd)
			wg.Done()
			fmt.Println("主机:", sshConfig.Address)
			fmt.Println(ss.Stdout)
		}(hostInfo)
		wg.Wait()

	}

}

// 这是方法用测试一些参数是否合格，比如检查grep是否存在，ip是否合法
func toTest() {

}
