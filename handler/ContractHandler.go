package handler

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetToAddress(client *ethclient.Client, contractAddr string, funcName string, abiJson []byte, params ...interface{}) (*common.Address, error) {
	// 解析abi
	contractABI, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	// data, _ := contractABI.Pack("getPair", common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"))
	data, _ := contractABI.Pack("getPair", params...)
	output, err := get(client, contractAddr, funcName, contractABI, data)
	if err != nil {
		return nil, err
	}

	var result *common.Address
	err = contractABI.UnpackIntoInterface(result, "getPair", output)
	if err != nil {
		log.Fatalf("failed to unpack result: %v", err)
	}
	log.Printf("call contract result: %v", result.String())
	return result, nil
}

func GetToBigInt(client *ethclient.Client, contractAddr string, funcName string, abiJson []byte) (*big.Int, error) {
	// 解析abi
	contractABI, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	data, _ := contractABI.Pack(funcName)
	output, err := get(client, contractAddr, funcName, contractABI, data)
	if err != nil {
		return nil, err
	}

	// 解析结果
	var result *big.Int
	err = contractABI.UnpackIntoInterface(&result, funcName, output)
	if err != nil {
		return nil, err
	}
	log.Printf("call contract result: %v", result.String())
	return result, nil
}

func get(client *ethclient.Client, contractAddr string, funcName string, contractABI abi.ABI, data []byte) ([]byte, error) {
	log.Printf("contract call, contractAddress: %s, funcName: %s", contractAddr, funcName)
	contractAddress := common.HexToAddress(contractAddr)

	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	// call
	output, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Printf("failed to call contract, contractAddress: %s, funcName: %s, err: %v", contractAddr, funcName, err)
		return nil, err
	}
	return output, nil
}
