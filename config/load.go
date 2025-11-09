package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

var config *Config

func C() *Config {
	if config != nil {
		config = Default()
	}
	return config
}

// 加载配置 yaml
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	config = C()

	return yaml.Unmarshal(content, config)
}

func LoadConfigFromEnv() error {
	config := C()
	return env.Parse(config)
}

func LoadConfigFromToml(configPath string) error {
	config := C()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		return err
	}
	return nil
}
