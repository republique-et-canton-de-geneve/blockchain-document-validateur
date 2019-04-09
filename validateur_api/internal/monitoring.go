package internal

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	//"time"
	"log"
)

type Sonde struct {
	ethereumActive bool
}

type TestStruct struct {
	Test string
}

type MonitoringEnv struct {
	NodeAddress   string
	LockedAddress string
	PrivateKey    string
}

var mn MonitoringEnv

func GetNodeSignal(ctx context.Context) bool {
	blkCtx, ok := BLKFromContext(ctx)
	if !ok {
		log.Fatalf("Could not obtain ClientConnector from context\n", ok)
		return false
	}
	//txHash:= common.HexToHash("75139f2e9f045987f67ab04541d03d7cd872e663b5efd758c20da42c89e652eb")
	//Above this comment is the line of code used for production version, the hash used is from the main net.
	//Below this comment is the line of code used for development version, the hash used is from the rinkeby testnet.
	txHash := common.HexToHash("d3851f8ee9bbd79a4cf332999a89a4b2c6b8d5c4c0c001ea85e95ab7997843c0")
	_, _, err := blkCtx.NC.GetTransaction(ctx, txHash)
	if err != nil {
		return false
	}
	return true
}

func InitMonitoring(nodeAddress string, lockedAddress string) MonitoringEnv {
	mn = MonitoringEnv{
		NodeAddress: nodeAddress,
		LockedAddress: lockedAddress,
	}
	return mn
}
