package config

type Mysql struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         string `mapstructure:"port" yaml:"port" json:"port"`
	Database     string `mapstructure:"database" yaml:"database" json:"database"`
	Username     string `mapstructure:"username" yaml:"username" json:"username"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" yaml:"maxOpenConns" json:"maxOpenConns"`
}
