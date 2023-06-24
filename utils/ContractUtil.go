package utils

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"interest-arbitrage/constant"
	dex2 "interest-arbitrage/contract/dex"
	c_erc20 "interest-arbitrage/contract/erc20"
	c_factory "interest-arbitrage/contract/factory"
	pair2 "interest-arbitrage/contract/pair"
	"interest-arbitrage/global"
	"interest-arbitrage/server"
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

func IsZeroAddress(address string) bool {
	return bytes.Equal([]byte(address), []byte(constant.ZeroAddress))
}

func BuildTransactOpts(privateKeys ...string) (*bind.TransactOpts, error) {
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

		opts, err = bind.NewKeyedTransactorWithChainID(auth, server.UsableNodeServer.ChainId)
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
