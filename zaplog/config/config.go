package config

type Config struct {
	Default DefaultOptions `yaml:"default"`
}

type DefaultOptions struct {
	Listen   int    `yaml:"listen"`
	LogType  string `yaml:"log_type"`
	LogDir   string `yaml:"log_dir"`
	LogLevel string `yaml:"log_level"`
}
