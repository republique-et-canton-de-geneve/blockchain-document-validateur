package internal

import (
	"context"
	//"fmt"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/ethclient"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	//"math"
	//"math/big"
)

type Sonde struct {
	ethereumActive					bool
}

type TestStruct struct {
	Test 								string
}

type MonitoringEnv struct {
	NodeAddress                        	string
	LockedAddress                       string
	PrivateKey							string
}

var mn MonitoringEnv

func GetNodeSignal(ctx context.Context) (bool){
	blkCtx, ok := BLKFromContext(ctx)
	if !ok {
		log.Println("Could not obtain ClientConnector from context\n")
		return false
	}
	// test connexion
	i, err := blkCtx.NC.SuggestGasPrice(ctx)
	if err != nil {
		return false
	}
	log.Println(i)
	return true
}


func InitMonitoring(nodeAddress string, lockedAddress string, privateKey string) MonitoringEnv {
	mn = MonitoringEnv{NodeAddress: nodeAddress,
		LockedAddress: lockedAddress,
		PrivateKey: privateKey,}
	return mn
}
