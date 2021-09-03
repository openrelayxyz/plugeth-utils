package core

import (
  "context"
  "math/big"
  "time"
  "github.com/holiman/uint256"
)

type Backend interface {
	// General Ethereum API
	Downloader() Downloader
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// ChainDb() Database
	// AccountManager() *accounts.Manager
	ExtRPCEnabled() bool
	RPCGasCap() uint64        // global gas cap for eth_call over rpc: DoS protection
	RPCTxFeeCap() float64     // global tx fee cap for all transaction related APIs
	UnprotectedAllowed() bool // allows only for EIP155 transactions.

	// Blockchain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, number int64) ([]byte, error) // RLP encoded header
	HeaderByHash(ctx context.Context, hash Hash) ([]byte, error)
	// HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error)
	CurrentHeader() []byte // RLP encoded header
	CurrentBlock() []byte // RLP encoded block
	BlockByNumber(ctx context.Context, number int64) ([]byte, error) // RLP encoded block
	BlockByHash(ctx context.Context, hash Hash) ([]byte, error) // RLP encoded block
	// BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error)
	// StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	// StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error)
	GetReceipts(ctx context.Context, hash Hash) ([]byte, error) // JSON encoded receipt (YES, JSON)
	GetTd(ctx context.Context, hash Hash) *big.Int

	SubscribeChainEvent(ch chan<- ChainEvent) Subscription
	SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) Subscription
	SubscribeChainSideEvent(ch chan<- ChainSideEvent) Subscription

	// Transaction pool API
	SendTx(ctx context.Context, signedTx []byte) error // RLP Encoded Transaction
	GetTransaction(ctx context.Context, txHash Hash) ([]byte, Hash, uint64, uint64, error) // RLP Encoded transaction
	GetPoolTransactions() ([][]byte, error) // []RLP ecnoded transactions
	GetPoolTransaction(txHash Hash) []byte // RLP encoded transaction
	GetPoolNonce(ctx context.Context, addr Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[Address][][]byte, map[Address][][]byte) // RLP encoded transactions
	SubscribeNewTxsEvent(chan<- NewTxsEvent) Subscription

	// Filter API
	BloomStatus() (uint64, uint64)
	GetLogs(ctx context.Context, blockHash Hash) ([][]byte, error) // []RLP encoded logs
	SubscribeLogsEvent(ch chan<- [][]byte) Subscription // []RLP encoded logs
	SubscribePendingLogsEvent(ch chan<- [][]byte) Subscription // RLP Encoded logs
	SubscribeRemovedLogsEvent(ch chan<- []byte) Subscription // RLP encoded logs

	// ChainConfig() *params.ChainConfig
	// Engine() consensus.Engine
}


type OpCode byte

type TracerResult interface {
	CaptureStart(from Address, to Address, create bool, input []byte, gas uint64, value *big.Int)
	CaptureState(pc uint64, op OpCode, gas, cost uint64, scope ScopeContext, rData []byte, depth int, err error)
	CaptureFault(pc uint64, op OpCode, gas, cost uint64, scope ScopeContext, depth int, err error)
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error)
	Result() interface{}
}

type ScopeContext interface {
	Memory()   Memory
	Stack()    Stack
	Contract() Contract
}

type Memory interface {
  GetCopy(int64, int64) []byte
  Len()     int
}

type Stack interface {
  Back(n int) *uint256.Int
  Len() int
}

type Contract interface {
  AsDelegate() Contract
  GetOp(n uint64) OpCode
  GetByte(n uint64) byte
  Caller() Address
  UseGas(gas uint64) (ok bool)
  Address() Address
  Value() *big.Int
}

type Downloader interface{
  Progress() Progress
}

type Progress interface{
  StartingBlock() uint64
  CurrentBlock() uint64
  HighestBlock() uint64
  PulledStates() uint64
  KnownStates() uint64
}


type Node interface{
  Server() Server
  DataDir() string
  InstanceDir() string
  IPCEndpoint() string
  HTTPEndpoint() string
  WSEndpoint() string
  ResolvePath(x string) string
}


type Server interface{
  PeerCount() int
}


type Logger interface {
  Trace(string, ...interface{})
  Debug(string, ...interface{})
  Info(string, ...interface{})
  Warn(string, ...interface{})
  Crit(string, ...interface{})
  Error(string, ...interface{})
}

type PluginLoader interface{
  Lookup(name string, validate func(interface{}) bool) []interface{}
}


