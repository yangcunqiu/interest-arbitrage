package server

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"interest-arbitrage/handler"
	"interest-arbitrage/model"
	"log"
	"math"
	"math/big"
	"time"
)

var UsablePriceServer *PriceServer

var uniswapV2FactoryAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
var uniswapV2FactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
var uniswapV2PairABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

type PriceServer struct {
	chainLinkConfig      []model.ChainLinkConfig
	uniswapV2PriceConfig []model.UniswapV2PriceConfig
	node                 *NodeServer
	stop                 chan struct{}
}

func GetPriceServer(node *NodeServer) *PriceServer {
	return &PriceServer{
		chainLinkConfig:      make([]model.ChainLinkConfig, 0),
		uniswapV2PriceConfig: make([]model.UniswapV2PriceConfig, 0),
		node:                 node,
		stop:                 make(chan struct{}),
	}
}

func (p *PriceServer) Stop() {
	p.stop <- struct{}{}
}

func (p *PriceServer) Restart() {
	p.Stop()
	p.chainLinkConfig = []model.ChainLinkConfig{}
	p.uniswapV2PriceConfig = []model.UniswapV2PriceConfig{}
	p.Start()
}

func (p *PriceServer) Start() {
	// 查询配置
	if len(p.chainLinkConfig) == 0 {
		p.chainLinkConfig = model.ListChainLinkConfig()
		p.uniswapV2PriceConfig = model.ListUniswapPriceConfig()
	}

	// 获取chainLink价格
	// for _, config := range p.chainLinkConfig {
	//	go getChainLinkPrice(p.node.client, config, p.stop)
	// }

	// 获取uniswap价格
	for _, config := range p.uniswapV2PriceConfig {
		go getUniswapPrice(p.node.client, &config, p.stop)
	}

}

func getUniswapPrice(client *ethclient.Client, config *model.UniswapV2PriceConfig, stop chan struct{}) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			uniswapPriceTask(client, config)
		case <-stop:
			ticker.Stop()
			return
		}
	}
}

func uniswapPriceTask(client *ethclient.Client, config *model.UniswapV2PriceConfig) {
	// 获取pair地址
	pairAddress := config.PairAddress
	if pairAddress == "" {
		address, err := handler.GetToAddress(client, uniswapV2FactoryAddress, "getPair", bytes.NewBufferString(uniswapV2FactoryABI).Bytes(),
			common.HexToAddress(config.Token0Address), common.HexToAddress(config.Token1Address))
		if err != nil {
			log.Printf("failed to get pair: %v", err)
			return
		}
		pairAddress = address.String()
		config.PairAddress = pairAddress
		model.UpdateUniswapPriceConfig(config)
	}

	price0CumulativeLast, err := handler.GetToBigInt(client, pairAddress, "price0CumulativeLast", bytes.NewBufferString(uniswapV2PairABI).Bytes())
	price1CumulativeLast, err := handler.GetToBigInt(client, pairAddress, "price1CumulativeLast", bytes.NewBufferString(uniswapV2PairABI).Bytes())
	if err != nil {
		log.Printf("failed to get cumulativeLast: %v", err)
		return
	}
	reserve, err := handler.GetReserve(client, pairAddress, bytes.NewBufferString(uniswapV2PairABI).Bytes())
	if err != nil {
		log.Printf("failed to get reserve: %v", err)
		return
	}

	var price0Average string
	var price1Average string
	// 查下上一次的
	last := model.GetTokenCumulativeLastByPairAddress(config.PairAddress)
	if last != nil {
		price0Average = last.Price0Average
		price1Average = last.Price1Average
		// 计算价格
		block := handler.GetBlockHeader(client)
		currentBlockTimestamp := uint(math.Mod(float64(block.Time), math.Pow(2, 32)))
		if currentBlockTimestamp != uint(reserve.BlockTimestamp) {
			b := fraction(reserve.Reserve0, reserve.Reserve1)
		}

		getCurrentCumulativePrices(block.Time, reserve.BlockTimestamp, price0CumulativeLast, price1CumulativeLast)

		timeElapsed := reserve.BlockTimestamp - last.BlockTimestampLast
		if timeElapsed > 0 {
			decimalsBigInt := new(big.Int)
			pow10 := math.Pow10(config.Decimals)
			decimals := decimalsBigInt.SetInt64(int64(pow10))
			// token0
			result0 := new(big.Int)
			timeElapsedBig := big.NewInt(int64(timeElapsed))

			bigInt0 := new(big.Int)
			price0ByPairAddress, _ := bigInt0.SetString(last.Price0CumulativeLast, 10)

			result0.Sub(price0CumulativeLast, price0ByPairAddress).Div(result0, timeElapsedBig).Div(result0, decimals)
			price0Average = result0.String()

			// token1
			result1 := new(big.Int)
			bigInt1 := new(big.Int)
			price1ByPairAddress, _ := bigInt1.SetString(last.Price1CumulativeLast, 10)

			result1.Sub(price1CumulativeLast, price1ByPairAddress).Div(result1, timeElapsedBig).Div(result1, decimals)
			price1Average = result1.String()
		}
	}
	// 保存
	tcl := &model.UniswapV2TokenCumulativeLast{
		PairAddress:          pairAddress,
		Price0Average:        price0Average,
		Price1Average:        price1Average,
		Reserve0:             reserve.Reserve0.String(),
		Reserve1:             reserve.Reserve1.String(),
		Price0CumulativeLast: price0CumulativeLast.String(),
		Price1CumulativeLast: price1CumulativeLast.String(),
		BlockTimestampLast:   reserve.BlockTimestamp,
	}
	model.SaveOrUpdateTokenCumulativeLast(tcl)

	// 解析价格

}

func fraction(numerator *big.Int, denominator *big.Int) *big.Int {
	result := new(big.Int)

	big1 := new(big.Int)
	big1.SetInt64(1)

	big144 := new(big.Int)
	big144.SetInt64(int64(math.Pow(2, 144))).Sub(big144, big1)

	big112 := new(big.Int)
	big112.SetInt64(int64(math.Pow(2, 112)))

	if numerator.Cmp(big144) <= 0 {
		result = numerator.Mul(numerator, big144).Div(numerator, denominator)
	} else {
		result = numerator.Mul(numerator, big112).Div(numerator, denominator)
	}
	return result
}

func getCurrentCumulativePrices(currentTimestamp uint64, reserveTimestamp uint32, price0Cumulative *big.Int, price1Cumulative *big.Int) {

}

func getChainLinkPrice(client *ethclient.Client, config model.ChainLinkConfig, stop chan struct{}) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ticker.C:
			chainLinkPriceTask(client, config)
		case <-stop:
			ticker.Stop()
			return
		}
	}
}

func chainLinkPriceTask(client *ethclient.Client, config model.ChainLinkConfig) {
	log.Println("getChainLinkPrice start ", config.PairName)
	price, err := handler.GetToBigInt(client, config.ContractAddr, config.PriceFuncName, bytes.NewBufferString(config.AbiJson).Bytes())
	if err != nil {
		log.Printf("failed to getChainLinkPrice: %v", err)
	}
	// 保存价格
	fPrice := new(big.Float)
	fPrice.SetString(price.String())
	readPrice := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(config.Decimals)))
	log.Printf("getChainLinkPrice success, pairName: %s, readPrice: %v", config.PairName, readPrice)

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
