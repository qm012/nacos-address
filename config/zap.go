package config

type Zap struct {
	Level     string `yaml:"level"`
	Filename  string `yaml:"filename"`
	MaxSize   int    `yaml:"max_size"`
	MaxBackup int    `yaml:"max_backup"`
	MaxAge    int    `yaml:"max_age"`
}
