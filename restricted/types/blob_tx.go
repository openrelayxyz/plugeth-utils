// Copyright 2023 The go-ethereum Authors
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
	"bytes"
	"math/big"
	"crypto/sha256"
	"fmt"

	"github.com/openrelayxyz/plugeth-utils/core"
	"github.com/openrelayxyz/plugeth-utils/restricted/params"
	"github.com/openrelayxyz/plugeth-utils/restricted/rlp"
	"github.com/openrelayxyz/plugeth-utils/restricted/crypto/kzg4844"
	"github.com/holiman/uint256"
)

// BlobTx represents an EIP-4844 transaction.
type BlobTx struct {
	ChainID    *uint256.Int
	Nonce      uint64
	GasTipCap  *uint256.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *uint256.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         core.Address
	Value      *uint256.Int
	Data       []byte
	AccessList AccessList
	BlobFeeCap *uint256.Int // a.k.a. maxFeePerBlobGas
	BlobHashes []core.Hash

	// A blob transaction can optionally contain blobs. This field must be set when BlobTx
	// is used to create a transaction for sigining.
	Sidecar *BlobTxSidecar `rlp:"-"`

	// Signature values
	V *uint256.Int `json:"v" gencodec:"required"`
	R *uint256.Int `json:"r" gencodec:"required"`
	S *uint256.Int `json:"s" gencodec:"required"`
}

// BlobTxSidecar contains the blobs of a blob transaction.
type BlobTxSidecar struct {
	Blobs       []kzg4844.Blob       // Blobs needed by the blob pool
	Commitments []kzg4844.Commitment // Commitments needed by the blob pool
	Proofs      []kzg4844.Proof      // Proofs needed by the blob pool
}

// BlobHashes computes the blob hashes of the given blobs.
func (sc *BlobTxSidecar) BlobHashes() []core.Hash {
	h := make([]core.Hash, len(sc.Commitments))
	for i := range sc.Blobs {
		h[i] = blobHash(&sc.Commitments[i])
	}
	return h
}

// encodedSize computes the RLP size of the sidecar elements. This does NOT return the
// encoded size of the BlobTxSidecar, it's just a helper for tx.Size().
func (sc *BlobTxSidecar) encodedSize() uint64 {
	var blobs, commitments, proofs uint64
	for i := range sc.Blobs {
		blobs += rlp.BytesSize(sc.Blobs[i][:])
	}
	for i := range sc.Commitments {
		commitments += rlp.BytesSize(sc.Commitments[i][:])
	}
	for i := range sc.Proofs {
		proofs += rlp.BytesSize(sc.Proofs[i][:])
	}
	return rlp.ListSize(blobs) + rlp.ListSize(commitments) + rlp.ListSize(proofs)
}

// blobTxWithBlobs is used for encoding of transactions when blobs are present.
type blobTxWithBlobs struct {
	BlobTx      *BlobTx
	Blobs       []kzg4844.Blob
	Commitments []kzg4844.Commitment
	Proofs      []kzg4844.Proof
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *BlobTx) copy() TxData {
	cpy := &BlobTx{
		Nonce: tx.Nonce,
		To:    tx.To,
		Data:  core.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		BlobHashes: make([]core.Hash, len(tx.BlobHashes)),
		Value:      new(uint256.Int),
		ChainID:    new(uint256.Int),
		GasTipCap:  new(uint256.Int),
		GasFeeCap:  new(uint256.Int),
		BlobFeeCap: new(uint256.Int),
		V:          new(uint256.Int),
		R:          new(uint256.Int),
		S:          new(uint256.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	copy(cpy.BlobHashes, tx.BlobHashes)

	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.GasTipCap)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.BlobFeeCap != nil {
		cpy.BlobFeeCap.Set(tx.BlobFeeCap)
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
	if tx.Sidecar != nil {
		cpy.Sidecar = &BlobTxSidecar{
			Blobs:       append([]kzg4844.Blob(nil), tx.Sidecar.Blobs...),
			Commitments: append([]kzg4844.Commitment(nil), tx.Sidecar.Commitments...),
			Proofs:      append([]kzg4844.Proof(nil), tx.Sidecar.Proofs...),
		}
	}
	return cpy
}

// accessors for innerTx.
func (tx *BlobTx) txType() byte           { return BlobTxType }
func (tx *BlobTx) chainID() *big.Int      { return tx.ChainID.ToBig() }
func (tx *BlobTx) accessList() AccessList { return tx.AccessList }
func (tx *BlobTx) data() []byte           { return tx.Data }
func (tx *BlobTx) gas() uint64            { return tx.Gas }
func (tx *BlobTx) gasFeeCap() *big.Int    { return tx.GasFeeCap.ToBig() }
func (tx *BlobTx) gasTipCap() *big.Int    { return tx.GasTipCap.ToBig() }
func (tx *BlobTx) gasPrice() *big.Int     { return tx.GasFeeCap.ToBig() }
func (tx *BlobTx) value() *big.Int        { return tx.Value.ToBig() }
func (tx *BlobTx) nonce() uint64          { return tx.Nonce }
func (tx *BlobTx) to() *core.Address    { tmp := tx.To; return &tmp }
func (tx *BlobTx) blobGas() uint64        { return params.BlobTxBlobGasPerBlob * uint64(len(tx.BlobHashes)) }
func (tx *BlobTx) blobGasFeeCap() *big.Int { return tx.BlobFeeCap.ToBig() }
func (tx *BlobTx) blobHashes() []core.Hash { return tx.BlobHashes }

func (tx *BlobTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap.ToBig())
	}
	tip := dst.Sub(tx.GasFeeCap.ToBig(), baseFee)
	if tip.Cmp(tx.GasTipCap.ToBig()) > 0 {
		tip.Set(tx.GasTipCap.ToBig())
	}
	return tip.Add(tip, baseFee)
}

func (tx *BlobTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V.ToBig(), tx.R.ToBig(), tx.S.ToBig()
}

func (tx *BlobTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID.SetFromBig(chainID)
	tx.V.SetFromBig(v)
	tx.R.SetFromBig(r)
	tx.S.SetFromBig(s)
}

func (tx *BlobTx) withoutSidecar() *BlobTx {
	cpy := *tx
	cpy.Sidecar = nil
	return &cpy
}

func (tx *BlobTx) encode(b *bytes.Buffer) error {
	if tx.Sidecar == nil {
		return rlp.Encode(b, tx)
	}
	inner := &blobTxWithBlobs{
		BlobTx:      tx,
		Blobs:       tx.Sidecar.Blobs,
		Commitments: tx.Sidecar.Commitments,
		Proofs:      tx.Sidecar.Proofs,
	}
	return rlp.Encode(b, inner)
}

func (tx *BlobTx) decode(input []byte) error {
	// Here we need to support two formats: the network protocol encoding of the tx (with
	// blobs) or the canonical encoding without blobs.
	//
	// The two encodings can be distinguished by checking whether the first element of the
	// input list is itself a list.

	outerList, _, err := rlp.SplitList(input)
	if err != nil {
		return err
	}
	firstElemKind, _, _, err := rlp.Split(outerList)
	if err != nil {
		return err
	}

	if firstElemKind != rlp.List {
		if err := rlp.DecodeBytes(input, tx); err != nil {
			// 0xf8b6053d850165a0bc00858bb2c97000825208944f56ffc63c28b72f79b02e91f11a4707bac4043c8080c0858bb2c97000f842a00155acead2ea1da3a7f6e9c743f4d81b1cd9f5a3382e657aa4025837e2b1651fa0019db7c3b9244969140c36c20cdcc71a821e2a995714c1ec2169f3ebd63c648280a00567aaceaed2303b577178069e5f4bfdd70d445e9eff8e63f9f0a33de98dd782a060e39ff4d29ce28325e3b48707bb8dd170ed3671428a62efcab9812555eb6db1
			return fmt.Errorf("Decode error: %v - %v", err, firstElemKind)
		} else {
			return nil
		}
	}
	// It's a tx with blobs.
	var inner blobTxWithBlobs
	if err := rlp.DecodeBytes(input, &inner); err != nil {
		return err
	}
	*tx = *inner.BlobTx
	tx.Sidecar = &BlobTxSidecar{
		Blobs:       inner.Blobs,
		Commitments: inner.Commitments,
		Proofs:      inner.Proofs,
	}
	return nil
}

func blobHash(commit *kzg4844.Commitment) core.Hash {
	hasher := sha256.New()
	hasher.Write(commit[:])
	var vhash core.Hash
	hasher.Sum(vhash[:0])
	vhash[0] = params.BlobTxHashVersion
	return vhash
}
