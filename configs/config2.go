package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config2 struct {
	Database struct {
		MySQL struct {
			Host      string `yaml:"host"`
			Port      int    `yaml:"port"`
			Username  string `yaml:"username"`
			Password  string `yaml:"password"`
			DBName    string `yaml:"dbname"`
			Charset   string `yaml:"charset"`
			ParseTime bool   `yaml:"parseTime"`
			Loc       string `yaml:"loc"`
		} `yaml:"mysql"`
	} `yaml:"database2"`
}

func LoadConfig2(path string) (*Config2, error) {
	// 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析YAML
	var cfg Config2
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &cfg, nil
}
