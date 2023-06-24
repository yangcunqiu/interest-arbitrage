package config

type Env struct {
	NodeUrl    string `yaml:"nodeUrl"`
	PrivateKey string `yaml:"privateKey"`
	Address    string `yaml:"address"`
}
