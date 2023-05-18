package server

import (
	"bytes"
	"github.com/ethereum/go-ethereum/ethclient"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"log"
	"math"
	"math/big"
)

type PriceServer struct {
	chainLinkConfig []model.ChainLinkConfig
	node            *NodeServer
}

func GetPriceServer(node *NodeServer) *PriceServer {
	return &PriceServer{
		node: node,
	}
}

func (p *PriceServer) Start() {
	// 查询配置
	if len(p.chainLinkConfig) == 0 {
		p.chainLinkConfig = model.ListChainLinkConfig()
	}

	for _, config := range p.chainLinkConfig {
		// 获取价格
		go getPrice(p.node.client, config)
	}
}

func getPrice(client *ethclient.Client, config model.ChainLinkConfig) {
	price, err := handler.GetToBigInt(client, config.ContractAddr, config.PriceFuncName, bytes.NewBufferString(config.AbiJson).Bytes())
	if err != nil {
		log.Printf("failed to getPrice: %v", err)
	}
	// 保存价格
	fPrice := new(big.Float)
	fPrice.SetString(price.String())
	readPrice := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(config.Decimals)))

	tp := model.TokenPrice{
		PairName:  config.PairName,
		Decimals:  config.Decimals,
		Price:     price.String(),
		ReadPrice: readPrice.String(),
		ChainId:   config.ChainId,
		Type:      config.Type,
	}
	model.SaveTokenPrice(&tp)
}
