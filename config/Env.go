package config

type Env struct {
	NodeUrl    string `yaml:"nodeUrl"`
	PrivateKey string `yaml:"privateKey"`
	Address    string `yaml:"address"`

	MTSNodeUrl string `yaml:"MTSNodeUrl"`
	ETHNodeUrl string `yaml:"ETHNodeUrl"`

	MTSPrivateKey  string `yaml:"MTSPrivateKey"`
	MTSPrivateKey1 string `yaml:"MTSPrivateKey1"`
	ETHPrivateKey  string `yaml:"ETHPrivateKey"`
	ETHPrivateKey1 string `yaml:"ETHPrivateKey1"`
}
