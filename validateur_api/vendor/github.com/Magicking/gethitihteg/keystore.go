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
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type AgentOperator struct {
	Address    common.Address
	privateKey *ecdsa.PrivateKey
	Transactor *bind.TransactOpts
}

func NewAgentOperator(privateKey *ecdsa.PrivateKey) (*AgentOperator, error) {
	if privateKey == nil {
		var err error
		if privateKey, err = ecdsa.GenerateKey(secp256k1.S256(), rand.Reader); err != nil {
			return nil, fmt.Errorf("NewAgentOperator, GenerateKey: %v", err)
		}
	}

	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	return &AgentOperator{
		Address:    addr,
		privateKey: privateKey,
		Transactor: &bind.TransactOpts{
			From: addr,
		},
	}, nil
}

func (ao *AgentOperator) Sign(rawTx *types.Transaction) (*types.Transaction, error) {
	transact_opt := bind.NewKeyedTransactor(ao.privateKey)
	signedTx, err := transact_opt.Signer(types.HomesteadSigner{}, ao.Transactor.From, rawTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
