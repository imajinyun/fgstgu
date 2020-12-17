package main

import (
	"e/lightsocks/cmd"
	"log"

	client "github.com/influxdata/influxdb1-client"
)

// DefaultListenAddr is default listen addr.
const DefaultListenAddr = ":7878"

var version = "client"

func main() {
	log.SetFlags(log.Lshortfile)

	config := &cmd.Config{ListenAddr: DefaultListenAddr}
	config.ReadConfig()
	config.SaveConfig()

	clt, err := client.NewClient(config.Password, config.Listen)

}
