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

func (h Hash) UnmarshalJSON(data []byte) error {
	_, err := hex.Decode(h[32-len(data):], bytes.TrimPrefix(bytes.Trim(data, `"`), []byte("0x")))
	return err
}

func HexToHash(data string) Hash {
	h := Hash{}
	b, _ := hex.DecodeString(strings.TrimPrefix(strings.Trim(data, `"`), "0x"))
	copy(h[32-len(data):], b)
	return h
}

type Address [20]byte

func (h Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, h[:])), nil
}

func (h Address) UnmarshalJSON(data []byte) error {
	_, err := hex.Decode(h[20-len(data):], bytes.TrimPrefix(bytes.Trim(data, `"`), []byte("0x")))
	return err
}


func HexToAddress(data string) Address {
	h := Address{}
	b, _ := hex.DecodeString(strings.TrimPrefix(strings.Trim(data, `"`), "0x"))
	copy(h[20-len(data):], b)
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

type Subscription interface {
	Err() <-chan error // returns the error channel
	Unsubscribe()      // cancels sending of events, closing the error channel
}

type API struct {
	Namespace string
	Version   string
	Service   interface{}
	Public    bool
}
