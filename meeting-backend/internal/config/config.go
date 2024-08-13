package config

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

var ConfigFilePath string = "config.toml"

func Config() *ConfigImpl {
	once.Do(func() {
		instance = loadConfig(ConfigFilePath)
	})

	return instance
}

func (c *ConfigImpl) ServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)
}

func (c *ConfigImpl) DatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.Charset,
	)
}

// ConfigImpl is a singleton that holds the configuration of the application.
var instance *ConfigImpl
var once sync.Once

func loadConfig(path string) *ConfigImpl {
	var conf ConfigImpl
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Println("加载配置文件“" + path + "”失败")
	}

	return &conf
}

type ConfigImpl struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	Charset  string `toml:"charset"`
}

type ServerConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}
