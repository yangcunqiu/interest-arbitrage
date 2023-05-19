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

func GetToBigInt(client *ethclient.Client, contractAddr string, funcName string, abiJson []byte) (*big.Int, error) {
	// 解析abi
	contractABI, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	output, err := get(client, contractAddr, funcName, contractABI)
	if err != nil {
		return nil, err
	}

	// 解析结果
	var result *big.Int
	err = contractABI.UnpackIntoInterface(&result, funcName, output)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func get(client *ethclient.Client, contractAddr string, funcName string, contractABI abi.ABI) ([]byte, error) {
	log.Printf("contract call, contractAddress: %s, funcName: %s", contractAddr, funcName)
	contractAddress := common.HexToAddress(contractAddr)
	data, _ := contractABI.Pack(funcName)
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
