package core

type Hash [32]byte
type Address [20]byte



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
