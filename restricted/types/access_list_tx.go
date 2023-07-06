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

//go:generate gencodec -type AccessTuple -out gen_access_tuple.go

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     core.Address `json:"address"        gencodec:"required"`
	StorageKeys []core.Hash  `json:"storageKeys"    gencodec:"required"`
}

// StorageKeys returns the total number of storage keys in the access list.
func (al AccessList) StorageKeys() int {
	sum := 0
	for _, tuple := range al {
		sum += len(tuple.StorageKeys)
	}
	return sum
}

// AccessListTx is the data of EIP-2930 access list transactions.
type AccessListTx struct {
	ChainIDv    *big.Int        // destination chain ID
	Noncev      uint64          // nonce of sender account
	GasPricev   *big.Int        // wei per gas
	Gasv        uint64          // gas limit
	Tov         *core.Address `rlp:"nil"` // nil means contract creation
	Valuev      *big.Int        // wei amount
	Datav       []byte          // contract invocation input data
	AccessListv AccessList      // EIP-2930 access list
	V, R, S    *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *AccessListTx) Copy() TxData {
	cpy := &AccessListTx{
		Noncev: tx.Noncev,
		Tov:    tx.Tov, // TODO: copy pointed-to address
		Datav:  core.CopyBytes(tx.Datav),
		Gasv:   tx.Gasv,
		// These are copied below.
		AccessListv: make(AccessList, len(tx.AccessListv)),
		Valuev:      new(big.Int),
		ChainIDv:    new(big.Int),
		GasPricev:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessListv, tx.AccessListv)
	if tx.Value != nil {
		cpy.Valuev.Set(tx.Valuev)
	}
	if tx.ChainID != nil {
		cpy.ChainIDv.Set(tx.ChainIDv)
	}
	if tx.GasPrice != nil {
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
func (tx *AccessListTx) TxType() byte           { return AccessListTxType }
func (tx *AccessListTx) ChainID() *big.Int      { return tx.ChainIDv }
func (tx *AccessListTx) Protected() bool        { return true }
func (tx *AccessListTx) AccessList() AccessList { return tx.AccessListv }
func (tx *AccessListTx) Data() []byte           { return tx.Datav }
func (tx *AccessListTx) Gas() uint64            { return tx.Gasv }
func (tx *AccessListTx) GasPrice() *big.Int     { return tx.GasPricev }
func (tx *AccessListTx) GasTipCap() *big.Int    { return tx.GasPricev }
func (tx *AccessListTx) GasFeeCap() *big.Int    { return tx.GasPricev }
func (tx *AccessListTx) Value() *big.Int        { return tx.Valuev }
func (tx *AccessListTx) Nonce() uint64          { return tx.Noncev }
func (tx *AccessListTx) To() *core.Address    { return tx.Tov }

func (tx *AccessListTx) RawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *AccessListTx) SetSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainIDv, tx.V, tx.R, tx.S = chainID, v, r, s
}
