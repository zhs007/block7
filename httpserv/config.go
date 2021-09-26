package block7serv

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config - configuration
type Config struct {
	BindAddr    string `yaml:"bindaddr"`
	IsDebugMode bool   `yaml:"isdebugmode"`
	DBPath      string `yaml:"dbpath"`
	DBEngine    string `yaml:"dbengine"`
	LogLevel    string `yaml:"loglevel"`
	LogPath     string `yaml:"logpath"`
	StatsToken  string `yaml:"statstoken"`
}

// LoadConfig - load config
func LoadConfig(fn string) (*Config, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.LogPath == "" {
		cfg.LogPath = "./logs"
	}

	if cfg.DBPath == "" {
		cfg.DBPath = "./data"
	}

	return cfg, nil
}
