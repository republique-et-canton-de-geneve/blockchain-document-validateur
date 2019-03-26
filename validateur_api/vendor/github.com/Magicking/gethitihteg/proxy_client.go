// Copyright 2017 Sylvain 6120 Laurent
// This file is part of the gethitihteg library.
//
// The gethitihteg library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gethitihteg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gethitihteg library. If not, see <http://www.gnu.org/licenses/>.

package blockchain

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NodeConnector struct {
	client  *ethclient.Client
	rawurl  string
	try_max int
}

func NewNodeConnector(rawurl string, try_max int) (*NodeConnector, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, fmt.Errorf("NewNodeConnector, Dial: %v", err)
	}
	return &NodeConnector{client: client, rawurl: rawurl, try_max: try_max}, nil
}

// Close close the RPC connection.
func (ec *NodeConnector) Close() {
	ec.Close()
}

// TODO Should return a fresh non-syncing ethereum client
func (nc *NodeConnector) GetClient(ctx context.Context) (*ethclient.Client, error) {
	i := 0
	for i < nc.try_max {
		sync_progress, err := nc.client.SyncProgress(ctx)
		if sync_progress == nil && err == nil {
			return nc.client, nil
		}
		if err != nil {
			client, err := ethclient.Dial(nc.rawurl)
			if err != nil {
				log.Printf("NewNodeConnector, try %d/%d, Dial: %v", i, nc.try_max, err)
				i++
				time.Sleep(30 * time.Second)
				continue
			}
			nc.client = client
		}
		sync_progress, err = nc.client.SyncProgress(ctx)
		if sync_progress != nil {
			i++
			log.Println("TODO: Loop here, wait(30s+ ?)/reconnect or abort ?")
			time.Sleep(30 * time.Second)
			continue
		}
		if err != nil {
			log.Printf("GetClient: %v", err)
			i++
			time.Sleep(30 * time.Second)
			continue
		}
		return nc.client, nil
	}
	return nil, fmt.Errorf("GetClient, too much try aborting")
}

func (nc *NodeConnector) WaitForFundOrCodeAt(account common.Address, blockNumber *big.Int, duration time.Duration) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	balance_ok := func() bool {
		balance, err := nc.GetBalanceAt(ctx, account, blockNumber)
		return !(err != nil || balance.Cmp(&big.Int{}) <= 0)
	}
	code_ok := func() bool {
		code, err := nc.CodeAt(ctx, account, blockNumber)
		return !(err != nil || len(code) == 0)
	}
	if balance_ok() || code_ok() {
		return true, nil
	}
	ticker := time.NewTicker(duration / 10)
	for {
		select {
		case <-ticker.C:
			if balance_ok() || code_ok() {
				return true, nil
			}
		case <-ctx.Done():
			ticker.Stop()
			return false, fmt.Errorf("WaitForFundAt: %v", ctx.Err())
		}
	}
	ticker.Stop() //Stop it anyway
	panic("WaitForFundAt: Dead code, should not go here")
	return false, fmt.Errorf("WaitForFundAt: Dead code, should not go here")
}

func (nc *NodeConnector) GetBalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	balance, err := nc.client.BalanceAt(ctx, account, blockNumber)
	if err != nil {
		log.Printf("BalanceAt, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("NodeConnector: %v", err)
		}
		balance, err = nc.client.BalanceAt(ctx, account, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("BalanceAt, re-try failed: %v", err)
		}
	}
	return balance, nil
}

func (nc *NodeConnector) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	code, err := nc.client.CodeAt(ctx, account, blockNumber)
	if err != nil {
		log.Printf("CodeAt, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("NodeConnector: %v", err)
		}
		code, err = nc.client.CodeAt(ctx, account, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("CodeAt, re-try failed: %v", err)
		}
	}
	return code, nil
}

func (nc *NodeConnector) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	output, err := nc.client.CallContract(ctx, call, blockNumber)
	if err != nil {
		log.Printf("CallContract, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("NodeConnector: %v", err)
		}
		output, err = nc.client.CallContract(ctx, call, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("CallContract, re-try failed: %v", err)
		}
	}
	return output, nil
}

func (nc *NodeConnector) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	//TODO CACHE ?
	code, err := nc.client.PendingCodeAt(ctx, account)
	if err != nil {
		log.Printf("PendingCodeAt, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("NodeConnector: %v", err)
		}
		code, err = nc.client.PendingCodeAt(ctx, account)
		if err != nil {
			return nil, fmt.Errorf("PendingCodeAt, re-try failed: %v", err)
		}
	}
	return code, nil
}

func (nc *NodeConnector) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	//TODO CACHE ?
	nonce, err := nc.client.PendingNonceAt(ctx, account)
	if err != nil {
		log.Printf("PendingNonceAt, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return 0, fmt.Errorf("NodeConnector: %v", err)
		}
		nonce, err = nc.client.PendingNonceAt(ctx, account)
		if err != nil {
			return 0, fmt.Errorf("PendingNonceAt, re-try failed: %v", err)
		}
	}
	return nonce, nil
}

func (nc *NodeConnector) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	price, err := nc.client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("SuggestGasPrice, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("NodeConnector: %v", err)
		}
		price, err = nc.client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, fmt.Errorf("SuggestGasPrice, re-try failed: %v", err)
		}
	}
	return price, nil
}

func (nc *NodeConnector) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	gas, err := nc.client.EstimateGas(ctx, msg)
	if err != nil {
		log.Printf("EstimateGas, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return 0, fmt.Errorf("NodeConnector: %v", err)
		}
		gas, err = nc.client.EstimateGas(ctx, msg)
		if err != nil {
			return 0, fmt.Errorf("EstimateGas, re-try failed: %v", err)
		}
	}
	return gas, nil
}

func (nc *NodeConnector) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := nc.client.SendTransaction(ctx, tx)
	if err != nil {
		log.Printf("SendTransaction, re-trying, error was: %v", err)
		_, err = nc.GetClient(ctx)
		if err != nil {
			return fmt.Errorf("NodeConnector: %v", err)
		}
		err = nc.client.SendTransaction(ctx, tx)
		if err != nil {
			return fmt.Errorf("SendTransaction, re-try failed: %v", err)
		}
	}
	return nil
}

// TransactionByHash returns the transaction with the given hash.
func (nc *NodeConnector) GetTransaction(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	tx, isPending, err = nc.client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}

type EventManager interface {
	FireEvent(ctx context.Context, evt *types.Log)
}

type ChanMsg int

const (
	DeathPill ChanMsg = iota
	Subscribe
)

func (nc *NodeConnector) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (nc *NodeConnector) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func (nc *NodeConnector) SubscribeToEvents(ctx context.Context, addr common.Address, evt_mgr EventManager) chan ChanMsg {
	cMsg := make(chan ChanMsg)

	go func() {
		q := ethereum.FilterQuery{Addresses: []common.Address{addr}}
		logChan := make(chan types.Log)
		defer close(logChan)
		for {
			select {
			case element := <-logChan:
				evt_mgr.FireEvent(ctx, &element)
			case msg := <-cMsg:
				switch msg {
				case Subscribe:
					for {
						_, err := nc.GetClient(ctx)
						if err != nil {
							log.Printf("Could not connect to Ethereum client: %v", err)
							time.Sleep(1 * time.Second) //TODO find a better way to yield
							continue
						}
						sub, err := nc.client.SubscribeFilterLogs(ctx, q, logChan)
						if err != nil {
							log.Printf("Could not subscribe to logs: %v", err)
							time.Sleep(1 * time.Second) //TODO find a better way to yield
							continue
						}
						go func() {
							log.Printf("SubscribeToEvents: %v", <-sub.Err())
							sub.Unsubscribe()
							time.Sleep(1 * time.Second) //TODO find a better way to yield
							cMsg <- Subscribe
						}()
						break
					}
				case DeathPill:
					return
				}
			}
		}
	}()

	cMsg <- Subscribe
	return cMsg
}
