package config

// System 系统配置
type System struct {
	// 端口
	Port int `mapstructure:"port" json:"port" yaml:"port"`
}
