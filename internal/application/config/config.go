package config

import "github.com/xiaodulala/admin-layout/internal/application/options"

// Config 配置结构,此架构配置与选项参数互通。
type Config struct {
	*options.Options
}

// CreateConfigFromOptions 选项结构转换为Config
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
