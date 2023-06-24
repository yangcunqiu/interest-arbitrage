package server

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var UsableNodeServer *NodeServer

type NodeServer struct {
	url     string
	Client  *ethclient.Client
	ChainId *big.Int
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
		log.Fatalf("failed to get ChainId: %v", err)
	}
	n.Client = client
	n.ChainId = chainID
}
