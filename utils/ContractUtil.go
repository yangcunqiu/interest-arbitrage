package utils

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"interest-arbitrage/constant"
	allot2 "interest-arbitrage/contract/allot"
	dex2 "interest-arbitrage/contract/dex"
	"interest-arbitrage/contract/eb"
	c_erc20 "interest-arbitrage/contract/erc20"
	c_factory "interest-arbitrage/contract/factory"
	"interest-arbitrage/contract/mb"
	pair2 "interest-arbitrage/contract/pair"
	"interest-arbitrage/contract/wmts"
	"interest-arbitrage/global"
	"interest-arbitrage/server"
	"math/big"
	"strings"
)

func GetERC20(erc20AddressStr string) (*c_erc20.Erc20, error) {
	erc20Address := common.HexToAddress(erc20AddressStr)
	erc20, err := c_erc20.NewErc20(erc20Address, server.UsableNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return erc20, nil
}

func GetPair(pairAddressStr string) (*pair2.Pair, error) {
	pairAddress := common.HexToAddress(pairAddressStr)
	pair, err := pair2.NewPair(pairAddress, server.UsableNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return pair, nil
}

func GetDexContract() (*dex2.Dex, error) {
	// 获取dex合约地址
	dexAddress := common.HexToAddress(global.ContractInfoMap[constant.Dex].Address)
	// 实例dex合约
	dex, err := dex2.NewDex(dexAddress, server.UsableNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return dex, nil
}

func GetFactoryContract() (*c_factory.Factory, error) {
	// 获取工厂合约地址
	factoryAddress := common.HexToAddress(global.ContractInfoMap[constant.Factory].Address)
	// 实例化工厂合约
	factory, err := c_factory.NewFactory(factoryAddress, server.UsableNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return factory, nil
}

func GetDexTokenAllot() (*allot2.Allot, error) {
	allotAddress := common.HexToAddress(global.ContractInfoMap[constant.ALLOT].Address)
	allot, err := allot2.NewAllot(allotAddress, server.UsableNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return allot, nil
}

func IsZeroAddress(address string) bool {
	return bytes.Equal([]byte(address), []byte(constant.ZeroAddress))
}

func BuildTransactOpts(chainId *big.Int, privateKeys ...string) (*bind.TransactOpts, error) {
	var opts *bind.TransactOpts
	// 签名
	if len(privateKeys) > 0 {
		privateKey := privateKeys[0]
		if strings.Index(privateKey, "0x") != -1 {
			privateKey = strings.TrimPrefix(privateKey, "0x")
		}
		// 签名
		auth, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, err
		}

		opts, err = bind.NewKeyedTransactorWithChainID(auth, chainId)
		if err != nil {
			return nil, err
		}

		// 获取gasPrice
		gasPrice, err := server.UsableNodeServer.Client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
		opts.GasPrice = gasPrice
	} else {
		opts = &bind.TransactOpts{}
	}
	return opts, nil
}

func GetETHBridge() (*eb.ETHBridge, error) {
	toAddress := common.HexToAddress(global.ContractInfoMap[constant.ETHBridge].Address)
	ethBridge, err := eb.NewETHBridge(toAddress, server.ETHNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return ethBridge, nil
}

func GetMTSBridge() (*mb.MTSBridge, error) {
	toAddress := common.HexToAddress(global.ContractInfoMap[constant.MTSBridge].Address)
	mtsBridge, err := mb.NewMTSBridge(toAddress, server.MTSNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return mtsBridge, nil
}

func GetWMTS() (*wmts.Wmts, error) {
	toAddress := common.HexToAddress(global.ContractInfoMap[constant.WMTS].Address)
	wmtsToken, err := wmts.NewWmts(toAddress, server.ETHNodeServer.Client)
	if err != nil {
		return nil, err
	}
	return wmtsToken, nil
}
