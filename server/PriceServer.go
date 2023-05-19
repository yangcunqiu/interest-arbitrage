package server

import (
	"bytes"
	"github.com/ethereum/go-ethereum/ethclient"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"log"
	"math"
	"math/big"
	"time"
)

var UsablePriceServer *PriceServer

type PriceServer struct {
	chainLinkConfig []model.ChainLinkConfig
	node            *NodeServer
	stop            chan struct{}
}

func GetPriceServer(node *NodeServer) *PriceServer {
	return &PriceServer{
		chainLinkConfig: make([]model.ChainLinkConfig, 0),
		node:            node,
		stop:            make(chan struct{}),
	}
}

func (p *PriceServer) Start() {
	// 查询配置
	if len(p.chainLinkConfig) == 0 {
		p.chainLinkConfig = model.ListChainLinkConfig()
	}

	for _, config := range p.chainLinkConfig {
		// 获取价格
		go getPrice(p.node.client, config, p.stop)
	}
}

func (p *PriceServer) Stop() {
	p.stop <- struct{}{}
}

func (p *PriceServer) Restart() {
	p.Stop()
	p.chainLinkConfig = []model.ChainLinkConfig{}
	p.Start()
}

func getPrice(client *ethclient.Client, config model.ChainLinkConfig, stop chan struct{}) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			priceTask(client, config)
		case <-stop:
			ticker.Stop()
			return
		}
	}
}

func priceTask(client *ethclient.Client, config model.ChainLinkConfig) {
	log.Println("getPrice start ", config.PairName)
	price, err := handler.GetToBigInt(client, config.ContractAddr, config.PriceFuncName, bytes.NewBufferString(config.AbiJson).Bytes())
	if err != nil {
		log.Printf("failed to getPrice: %v", err)
	}
	// 保存价格
	fPrice := new(big.Float)
	fPrice.SetString(price.String())
	readPrice := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(config.Decimals)))
	log.Printf("getPrice success, pairName: %s, readPrice: %v", config.PairName, readPrice)

	tp := model.TokenPrice{
		PairName:  config.PairName,
		Decimals:  config.Decimals,
		Price:     price.String(),
		ReadPrice: readPrice.String(),
		ChainId:   config.ChainId,
		Type:      config.Type,
	}
	model.SaveOrUpdateTokenPrice(&tp)
}
