package config

type Mysql struct {
	Ip               string `yaml:"ip"`
	Port             uint   `yaml:"port"`
	Database         string `yaml:"database"`
	Params           string `yaml:"params"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	MaxIdleConnCount int    `yaml:"maxIdleConnCount"`
	MaxOpenConnCount int    `yaml:"maxOpenConnCount"`
	ConnMaxLifetime  int    `yaml:"connMaxLifetime"`
}
