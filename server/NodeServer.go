package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type NodeServer struct {
	url    string
	client *ethclient.Client
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
	n.client = client
}
