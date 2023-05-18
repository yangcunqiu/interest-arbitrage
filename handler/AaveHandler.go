package handler

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

func GetPrice() {
	url := "https://eth-mainnet.g.alchemy.com/v2/0bPOYzslxoTx3U4jQRGNFFFo2mz5UgbO"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to %s, err: %v", url, err)
	}

	ethUsd := common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419")

	file, _ := os.ReadFile("C:\\project\\go\\interest-arbitrage\\abi\\EthUsdABi.json")

	contractABI, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	data, _ := contractABI.Pack("latestAnswer")
	callMsg := ethereum.CallMsg{
		To:   &ethUsd,
		Data: data,
	}
	output, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatalf("failed to call contract, err: %v", err)
	}

	var result *big.Int
	err = contractABI.UnpackIntoInterface(&result, "latestAnswer", output)
	if err != nil {
		log.Fatalf("failed to unpack: %v", err)
	}

	log.Println(result)
	log.Println(result.Div(result, big.NewInt(100000000)))
}
