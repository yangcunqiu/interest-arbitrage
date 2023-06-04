package server

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type NodeServer struct {
	url     string
	client  *ethclient.Client
	chainId *big.Int
}

func GetDefaultNodeServer(url string) *NodeServer {
	return &NodeServer{
		url: url,
	}
}

func (n *NodeServer) Start() {
	client, err := ethclient.Dial(n.url)
	if err != nil {
		log.Fatalf("failed to connect to %s, err: %v", n.url, err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chainId: %v", err)
	}
	n.client = client
	n.chainId = chainID
}
