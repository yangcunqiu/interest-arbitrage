package constant

import "interest-arbitrage/model"

var ContractInfos = make([]*model.ContractInfo, 0)

const ZeroAddress = "0x0000000000000000000000000000000000000000"

const (
	Factory = "PairFactory"
	Dex     = "MyDex"
	HL      = "HL"
	ML      = "ML"
	MDT     = "MDT"
	ALLOT   = "ALLOT"
)

var ContractNames = []string{
	"PairFactory",
	"MyDex",
	"HL",
	"ML",
	"MDT",
	"ALLOT",
}

// ContractAddress local
var ContractAddress = []string{
	"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
	"0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
	"0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512",
	"0x5FbDB2315678afecb367f032d93F642f64180aa3",
	"0xa82fF9aFd8f496c3d6ac40E2a0F282E47488CFc9",
	"0x1613beB3B2C4f22Ee086B2b38C1476A3cE7f78E8",
}

// ContractAddress aliyun
//
//	var ContractAddress = []string{
//		"0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
//		"0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
//		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
//		"0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
//	}
var ContractABIPaths = []string{
	"abi/PairFactory.json",
	"abi/MyDex.json",
	"abi/HL.json",
	"abi/ML.json",
	"abi/MDT.json",
	"abi/MyDexTokenAllot.json",
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
