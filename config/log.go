package config

// Log 日志配置
type Log struct {
	Level            string `mapstructure:"level" json:"level" yaml:"level"`
	Format           string `mapstructure:"format" json:"format" yaml:"format"`
	Dir              string `mapstructure:"dir" json:"dir" yaml:"dir"`
	MaxRetentionDays int    `mapstructure:"max_retention_days" json:"max_retention_days" yaml:"max_retention_days"`
	ShowLine         bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	ConsoleOutput    bool   `mapstructure:"console_output" json:"console_output" yaml:"console_output"`
}
