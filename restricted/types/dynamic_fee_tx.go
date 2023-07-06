// Copyright 2021 The go-ethereum Authors
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

type DynamicFeeTx struct {
	ChainIDv    *big.Int
	Noncev      uint64
	GasTipCapv  *big.Int
	GasFeeCapv  *big.Int
	Gasv        uint64
	Tov         *core.Address `rlp:"nil"` // nil means contract creation
	Valuev      *big.Int
	Datav       []byte
	AccessListv AccessList

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *DynamicFeeTx) Copy() TxData {
	cpy := &DynamicFeeTx{
		Noncev: tx.Noncev,
		Tov:    tx.Tov, // TODO: copy pointed-to address
		Datav:  core.CopyBytes(tx.Datav),
		Gasv:   tx.Gasv,
		// These are copied below.
		AccessListv: make(AccessList, len(tx.AccessListv)),
		Valuev:      new(big.Int),
		ChainIDv:    new(big.Int),
		GasTipCapv:  new(big.Int),
		GasFeeCapv:  new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessListv, tx.AccessListv)
	if tx.Valuev != nil {
		cpy.Valuev.Set(tx.Valuev)
	}
	if tx.ChainIDv != nil {
		cpy.ChainIDv.Set(tx.ChainIDv)
	}
	if tx.GasTipCapv != nil {
		cpy.GasTipCapv.Set(tx.GasTipCapv)
	}
	if tx.GasFeeCapv != nil {
		cpy.GasFeeCapv.Set(tx.GasFeeCapv)
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
func (tx *DynamicFeeTx) TxType() byte           { return DynamicFeeTxType }
func (tx *DynamicFeeTx) ChainID() *big.Int      { return tx.ChainIDv }
func (tx *DynamicFeeTx) Protected() bool        { return true }
func (tx *DynamicFeeTx) AccessList() AccessList { return tx.AccessListv }
func (tx *DynamicFeeTx) Data() []byte           { return tx.Datav }
func (tx *DynamicFeeTx) Gas() uint64            { return tx.Gasv }
func (tx *DynamicFeeTx) GasFeeCap() *big.Int    { return tx.GasFeeCapv }
func (tx *DynamicFeeTx) GasTipCap() *big.Int    { return tx.GasTipCapv }
func (tx *DynamicFeeTx) GasPrice() *big.Int     { return tx.GasFeeCapv }
func (tx *DynamicFeeTx) Value() *big.Int        { return tx.Valuev }
func (tx *DynamicFeeTx) Nonce() uint64          { return tx.Noncev }
func (tx *DynamicFeeTx) To() *core.Address    { return tx.Tov }

func (tx *DynamicFeeTx) RawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *DynamicFeeTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainIDv, tx.V, tx.R, tx.S = chainID, v, r, s
}
