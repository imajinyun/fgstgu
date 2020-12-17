package main

import (
	"e/lightsocks"
	"e/lightsocks/cmd"
	"e/lightsocks/server"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/phayes/freeport"
)

var version = "server"

func main() {
	log.SetFlags(log.Lshortfile)

	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))
	if err != nil {
		port, err = freeport.GetFreePort()
	}
	if err != nil {
		port = 7878
	}
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		Password:   lightsocks.RandPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	srv, err := server.NewServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(srv.Listen(func(addr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
lightsocks-server: %s 启动成功，配置如下：
服务器地址：%s
密码：%s
		`, version, addr, config.Password))
	}))
}
