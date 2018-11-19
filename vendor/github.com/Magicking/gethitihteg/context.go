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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewBlockchainContext(raw_url, pkey_hex string, retry int) (*BlockchainContext, error) {
	NC, err := NewNodeConnector(raw_url, retry)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize connection to node: %v", err)
	}
	pkey := common.FromHex(pkey_hex)
	privateKey, err := crypto.ToECDSA(pkey)
	if err != nil {
		return nil, fmt.Errorf("Failed to import PrivateKey: %v", err)
	}
	AO, err := NewAgentOperator(privateKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize AgentOperator: %v", err)
	}
	return &BlockchainContext{NC: NC, AO: AO}, nil
}
