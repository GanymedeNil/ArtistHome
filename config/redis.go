package config

type Redis struct {
	Host     string `mapstructure:"host" yaml:"host" json:"host"`
	Port     string `mapstructure:"port" yaml:"port" json:"port"`
	DB       string `mapstructure:"db" yaml:"db" json:"db"`
	Username string `mapstructure:"username" yaml:"username" json:"username"`
	Password string `mapstructure:"password" yaml:"password" json:"password"`
}
