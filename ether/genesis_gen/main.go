// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

type CliqueConfig struct {
	ChainID           int
	SignerAccounts    []string
	PrefondedAccounts []string
	Period            int
	Epoch             int
}

// TODO check CliqueConfig elements if exists
// TODO Other network parameters could be added
func (cc CliqueConfig) generateGenesis() []byte {
	// create slices

	var signers []common.Address
	var prefondeds []common.Address

	// OPT slicelarin boyu onceden hesaplanabilir
	for i, _ := range cc.SignerAccounts {
		signers = append(signers, common.HexToAddress(cc.SignerAccounts[i]))
	}

	for i, _ := range cc.PrefondedAccounts {
		prefondeds = append(prefondeds, common.HexToAddress(cc.PrefondedAccounts[i]))
	}

	fmt.Println(signers)
	fmt.Println(prefondeds)

	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(524288),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock:      big.NewInt(0),
			EIP150Block:         big.NewInt(0),
			EIP155Block:         big.NewInt(0),
			EIP158Block:         big.NewInt(0),
			ByzantiumBlock:      big.NewInt(0),
			ConstantinopleBlock: big.NewInt(0),
			PetersburgBlock:     big.NewInt(0),
			IstanbulBlock:       big.NewInt(0),
		},
	}

	genesis.Config.Ethash = new(params.EthashConfig)
	genesis.ExtraData = make([]byte, 32)

	genesis.Difficulty = big.NewInt(1)
	genesis.Config.Clique = &params.CliqueConfig{
		Period: 15,
		Epoch:  30000,
	}
	genesis.Config.Clique.Period = uint64(15)

	// Sort the signers and embed into the extra-data section
	for i := 0; i < len(signers); i++ {
		for j := i + 1; j < len(signers); j++ {
			if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
				signers[i], signers[j] = signers[j], signers[i]
			}
		}
	}
	genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)

	// Add signers
	for i, signer := range signers {
		copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
	}

	// Prefonded addresses
	for _, profunded := range prefondeds {
		// Read the address of the account to fund
		genesis.Alloc[profunded] = core.GenesisAccount{
			Balance: new(big.Int).Lsh(big.NewInt(1), 256-7), // 2^256 / 128 (allow many pre-funds without balance overflows)
		}
	}

	genesis.Config.ChainID = new(big.Int).SetUint64(uint64(cc.ChainID))

	jsonStr, _ := json.MarshalIndent(genesis, "", "  ")
	return jsonStr
}

// TODO
func addressHexStringToAddress(adressHex string) common.Address {
	recAddress := common.HexToAddress(adressHex)
	return recAddress
}

func main() {
	// key, _ := crypto.GenerateKey()
	// address := crypto.PubkeyToAddress(key.PublicKey)
	// addressString := address.Hex()

	// recAddr := addressHexStringToAddress(addressString)
	// if address == recAddr {
	// 	fmt.Println("Address recovered !")
	// }

	acc1 := "0x714D9c3B5e9C479059B89a70aAea4A3173e4597F"
	acc2 := "0x9Ee6056ca96243D54D0c9c07726F804C6179610D"

	chaindId := 29
	signers := []string{acc1, acc2}
	prefondeds := []string{acc1, acc2}
	epoch := 3000
	period := 15
	cliqueConfig := CliqueConfig{chaindId, signers, prefondeds, period, epoch}

	genesis := cliqueConfig.generateGenesis()
	fmt.Println(string(genesis))
}
