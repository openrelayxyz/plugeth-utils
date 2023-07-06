// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
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

package types

import (
	"math/big"

	"github.com/openrelayxyz/plugeth-utils/core"
)

// LegacyTx is the transaction data of regular Ethereum transactions.
type LegacyTx struct {
	Noncev    uint64          // nonce of sender account
	GasPricev *big.Int        // wei per gas
	Gasv      uint64          // gas limit
	Tov       *core.Address `rlp:"nil"` // nil means contract creation
	Valuev    *big.Int        // wei amount
	Datav     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}

// NewTransaction creates an unsigned legacy transaction.
// Deprecated: use NewTx instead.
func NewTransaction(nonce uint64, to core.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	return NewTx(&LegacyTx{
		Noncev:    nonce,
		Tov:       &to,
		Valuev:    amount,
		Gasv:      gasLimit,
		GasPricev: gasPrice,
		Datav:     data,
	})
}

// NewContractCreation creates an unsigned legacy transaction.
// Deprecated: use NewTx instead.
func NewContractCreation(nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction {
	return NewTx(&LegacyTx{
		Noncev:    nonce,
		Valuev:    amount,
		Gasv:      gasLimit,
		GasPricev: gasPrice,
		Datav:     data,
	})
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *LegacyTx) Copy() TxData {
	cpy := &LegacyTx{
		Noncev: tx.Noncev,
		Tov:    tx.Tov, // TODO: copy pointed-to address
		Datav:  core.CopyBytes(tx.Datav),
		Gasv:   tx.Gasv,
		// These are initialized below.
		Valuev:    new(big.Int),
		GasPricev: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Valuev != nil {
		cpy.Valuev.Set(tx.Valuev)
	}
	if tx.GasPricev != nil {
		cpy.GasPricev.Set(tx.GasPricev)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *LegacyTx) TxType() byte           { return LegacyTxType }
func (tx *LegacyTx) ChainID() *big.Int      { return deriveChainId(tx.V) }
func (tx *LegacyTx) AccessList() AccessList { return nil }
func (tx *LegacyTx) Data() []byte           { return tx.Datav }
func (tx *LegacyTx) Gas() uint64            { return tx.Gasv }
func (tx *LegacyTx) GasPrice() *big.Int     { return tx.GasPricev }
func (tx *LegacyTx) GasTipCap() *big.Int    { return tx.GasPricev }
func (tx *LegacyTx) GasFeeCap() *big.Int    { return tx.GasPricev }
func (tx *LegacyTx) Value() *big.Int        { return tx.Valuev }
func (tx *LegacyTx) Nonce() uint64          { return tx.Noncev }
func (tx *LegacyTx) To() *core.Address    { return tx.Tov }

func (tx *LegacyTx) RawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *LegacyTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}
