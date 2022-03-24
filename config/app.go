package config

type App struct {
	Env   string `mapstructure:"env" yaml:"env" json:"env"`
	Debug bool   `mapstructure:"debug" yaml:"debug" json:"debug"`
	Addr  string `mapstructure:"addr" yaml:"addr" json:"addr"`
}
