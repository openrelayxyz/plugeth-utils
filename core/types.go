package core

import (
	"strings"
	"bytes"
	"fmt"
	"encoding/hex"
)

type Hash [32]byte

func (h Hash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, h[:])), nil
}

func (h *Hash) UnmarshalJSON(data []byte) error {
	d := bytes.TrimPrefix(bytes.Trim(data, `"`), []byte("0x"))
	_, err := hex.Decode(h[(64 - len(d)) / 2:], d)
	return err
}

func (h Hash) Bytes() []byte {
	return ([]byte)(h[:])
}

func (h Hash) String() string {
	return fmt.Sprintf("%#x", h[:])
}

func HexToHash(data string) Hash {
	h := Hash{}
	b, _ := hex.DecodeString(strings.TrimPrefix(strings.Trim(data, `"`), "0x"))
	copy(h[32 - len(b):], b)
	return h
}

func BytesToHash(b []byte) Hash {
	h := Hash{}
	copy(h[32-len(b):], b)
	return h
}

type Address [20]byte

func (h Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, h[:])), nil
}

func (h *Address) UnmarshalJSON(data []byte) error {
	d := bytes.TrimPrefix(bytes.Trim(data, `"`), []byte("0x"))
	_, err := hex.Decode(h[(40 - len(d))/2:], d)
	return err
}

func (h Address) String() string {
	return fmt.Sprintf("%#x", h[:])
}


func HexToAddress(data string) Address {
	h := Address{}
	b, _ := hex.DecodeString(strings.TrimPrefix(strings.Trim(data, `"`), "0x"))
	copy(h[20 - len(b):], b)
	return h
}

func BytesToAddress(b []byte) Address {
	h := Address{}
	copy(h[20-len(b):], b)
	return h
}


type ChainEvent struct {
	Block []byte // RLP Encoded block
	Hash  Hash
	Logs  []byte // RLP encoded logs
}

type ChainSideEvent struct {
	Block []byte // RLP Encoded block
}

type ChainHeadEvent struct{
  Block []byte // RLP Encoded block
}

type NewTxsEvent struct{
  Txs [][]byte // []RLP encoded transaction
}

type API struct {
	Namespace string
	Version   string
	Service   interface{}
	Public    bool
}


func CopyBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b[:], a[:])
	return b
}
