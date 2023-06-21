package test

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

const (
	NodeUrl         = "http://127.0.0.1:8545/"
	FACTORY_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	MY_DEX_ADDRESS  = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
)

var client *ethclient.Client = nil

func initNode() {
	cl, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Fatalf("failed to connect to %s, err: %s", NodeUrl, err)
	}
	client = cl
}

// factory
func TestPair(t *testing.T) {
	initNode()

}

func TestQuote(t *testing.T) {
	initNode()

}

func TestAddLiquidity(t *testing.T) {
	initNode()

}
