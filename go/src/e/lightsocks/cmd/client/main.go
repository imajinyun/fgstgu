package main

import (
	"e/lightsocks/client"
	"e/lightsocks/cmd"
	"fmt"
	"log"
	"net"
)

// DefaultListenAddr is default listen addr.
const DefaultListenAddr = ":7878"

var version = "client"

func main() {
	log.SetFlags(log.Lshortfile)

	config := &cmd.Config{ListenAddr: DefaultListenAddr}
	config.ReadConfig()
	config.SaveConfig()

	clt, err := client.NewClient(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(clt.Listen(func(addr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
lightsocks-client: %s 启动成功，配置如下：
本地监听地址：%s
远程服务地址：%s
密码：%s
		`, version, addr, config.RemoteAddr, config.Password))
	}))
}
