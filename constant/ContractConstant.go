package constant

import "interest-arbitrage/model"

var ContractInfos = make([]*model.ContractInfo, 0)

const (
	Factory = "PairFactory"
	Dex     = "MyDex"
	HL      = "HL"
	ML      = "ML"
)

var ContractNames = []string{
	"PairFactory",
	"MyDex",
	"HL",
	"ML",
}
var ContractAddress = []string{
	"0x5FbDB2315678afecb367f032d93F642f64180aa3",
	"0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
	"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
	"0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
}
var ContractABIPaths = []string{
	"abi/PairFactory.json",
	"abi/MyDex.json",
	"abi/HL.json",
	"abi/ML.json",
}

func init() {
	for i, name := range ContractNames {
		info := &model.ContractInfo{
			Name:    name,
			Address: ContractAddress[i],
			ABIPath: ContractABIPaths[i],
		}
		ContractInfos = append(ContractInfos, info)
	}
}
