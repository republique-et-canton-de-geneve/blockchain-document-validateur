// Copyright 2017 Sylvain 6120 Laurent
// This file is part of the gethitihteg library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package blockchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateContractHelper(client bind.ContractBackend, hexKey, rawABI, contractBin string, cacheNonce *big.Int, params ...interface{}) (common.Address, *types.Transaction, *BoundContract, error) {
	key, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = cacheNonce
	contractABI, err := abi.JSON(strings.NewReader(rawABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	rawContract := common.FromHex(contractBin)
	addr, tx, c, err := DeployContract(auth, contractABI, rawContract, client, params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return addr, tx, c, nil
}
