package config

type Server struct {
	App   App   `mapstructure:"app" yaml:"app" json:"app"`
	Zap   Zap   `mapstructure:"zap" yaml:"zap" json:"zap"`
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
}
