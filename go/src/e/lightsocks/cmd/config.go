package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

var configPath string

// Config struct.
type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func init() {
	home, _ := homedir.Dir()
	configFilename := ".lightsocks.json"
	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}
	configPath = path.Join(home, configFilename)
}

// SaveConfig methof for Config.
func (c *Config) SaveConfig() {
	configJSON, _ := json.MarshalIndent(c, "", "  ")
	if err := ioutil.WriteFile(configPath, configJSON, 0644); err != nil {
		log.Fatalf("保存配置到文件 %s 出错：%s", configPath, err)
	}
	log.Printf("保存配置到文件 %s 成功\n", configPath)
}

// ReadConfig method for Config.
func (c *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("从文件 %s 中读取配置\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("打开配置文件 %s 出错：%s", configPath, err)
		}
		defer file.Close()
		if err = json.NewDecoder(file).Decode(c); err != nil {
			log.Fatalf("格式不合法的 JSON 配置文件：%s\n", file.Name())
		}
		if err != nil {

		}
	}
}
