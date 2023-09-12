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
	CurrentHeader() []byte                                           // RLP encoded header
	CurrentBlock() []byte                                            // RLP encoded block
	BlockByNumber(ctx context.Context, number int64) ([]byte, error) // RLP encoded block
	BlockByHash(ctx context.Context, hash Hash) ([]byte, error)      // RLP encoded block
	// BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error)
	// StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error)
	// StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error)
	GetReceipts(ctx context.Context, hash Hash) ([]byte, error) // JSON encoded receipt (YES, JSON)
	GetTd(ctx context.Context, hash Hash) *big.Int

	SubscribeChainEvent(ch chan<- ChainEvent) Subscription
	SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) Subscription
	SubscribeChainSideEvent(ch chan<- ChainSideEvent) Subscription

	// Transaction pool API
	SendTx(ctx context.Context, signedTx []byte) error                                     // RLP Encoded Transaction
	GetTransaction(ctx context.Context, txHash Hash) ([]byte, Hash, uint64, uint64, error) // RLP Encoded transaction
	GetPoolTransactions() ([][]byte, error)                                                // []RLP ecnoded transactions
	GetPoolTransaction(txHash Hash) []byte                                                 // RLP encoded transaction
	GetPoolNonce(ctx context.Context, addr Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[Address][][]byte, map[Address][][]byte) // RLP encoded transactions
	SubscribeNewTxsEvent(chan<- NewTxsEvent) Subscription

	// Filter API
	BloomStatus() (uint64, uint64)
	GetLogs(ctx context.Context, blockHash Hash) ([][]byte, error) // []RLP encoded logs
	SubscribeLogsEvent(ch chan<- [][]byte) Subscription            // []RLP encoded logs
	SubscribePendingLogsEvent(ch chan<- [][]byte) Subscription     // RLP Encoded logs
	SubscribeRemovedLogsEvent(ch chan<- []byte) Subscription       // RLP encoded logs

	GetTrie(hash Hash) (Trie, error)
	GetAccountTrie(stateRoot Hash, account Address) (Trie, error)
	GetContractCode(Hash) ([]byte, error)

	// ChainConfig() *params.ChainConfig
	// Engine() consensus.Engine
}

type OpCode byte

type BlockTracer interface {
	TracerResult
	PreProcessBlock(hash Hash, number uint64, encoded []byte)
	PreProcessTransaction(tx Hash, block Hash, i int)
	BlockProcessingError(tx Hash, block Hash, err error)
	PostProcessTransaction(tx Hash, block Hash, i int, receipt []byte)
	PostProcessBlock(block Hash)
}

// The implementation of CaptureEnd below diverges from foundation Geth, we pass dummy variables in PluGeth
// in order to preserve the implementation of the tracing plugins in Plugeth-Plugins.
type TracerResult interface {
	CaptureStart(from Address, to Address, create bool, input []byte, gas uint64, value *big.Int)
	CaptureState(pc uint64, op OpCode, gas, cost uint64, scope ScopeContext, rData []byte, depth int, err error)
	CaptureFault(pc uint64, op OpCode, gas, cost uint64, scope ScopeContext, depth int, err error)
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error)
	CaptureEnter(typ OpCode, from Address, to Address, input []byte, gas uint64, value *big.Int)
	CaptureExit(output []byte, gasUsed uint64, err error)
	Result() (interface{}, error)
}

type PreTracer interface {
	CapturePreStart(from Address, to *Address, input []byte, gas uint64, value *big.Int)
}

type StateDB interface {
	GetBalance(Address) *big.Int

	GetNonce(Address) uint64

	GetCodeHash(Address) Hash
	GetCode(Address) []byte
	GetCodeSize(Address) int

	GetRefund() uint64

	GetCommittedState(Address, Hash) Hash
	GetState(Address, Hash) Hash

	HasSuicided(Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(Address) bool

	AddressInAccessList(addr Address) bool
	SlotInAccessList(addr Address, slot Hash) (addressOk bool, slotOk bool)

	IntermediateRoot(deleteEmptyObjects bool) Hash
}

type RWStateDB interface {
	StateDB
}

type ScopeContext interface {
	Memory() Memory
	Stack() Stack
	Contract() Contract
}

type Memory interface {
	GetCopy(int64, int64) []byte
	Len() int
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
	Address() Address
	Value() *big.Int
	Input() []byte
	Code() []byte
}

type Downloader interface {
	Progress() Progress
}

type Progress interface {
	StartingBlock() uint64
	CurrentBlock() uint64
	HighestBlock() uint64
	PulledStates() uint64
	KnownStates() uint64
	SyncedAccounts() uint64
	SyncedAccountBytes() uint64
	SyncedBytecodes() uint64
	SyncedBytecodeBytes() uint64
	SyncedStorage() uint64
	SyncedStorageBytes() uint64
	HealedTrienodes() uint64
	HealedTrienodeBytes() uint64
	HealedBytecodes() uint64
	HealedBytecodeBytes() uint64
	HealingTrienodes() uint64
	HealingBytecode() uint64
}

type Node interface {
	Server() Server
	DataDir() string
	InstanceDir() string
	IPCEndpoint() string
	HTTPEndpoint() string
	WSEndpoint() string
	ResolvePath(x string) string
	Attach() (Client, error)
	Close() error
}

type Client interface {
	Call(interface{}, string, ...interface{}) error
}

type Server interface {
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

type PluginLoader interface {
	Lookup(name string, validate func(interface{}) bool) []interface{}
	GetFeed() Feed
}

type Subscription interface {
	Err() <-chan error // returns the error channel
	Unsubscribe()      // cancels sending of events, closing the error channel
}

type Feed interface {
	Send(interface{}) int
	Subscribe(channel interface{}) Subscription
}

type BlockContext struct {
	Coinbase    Address
	GasLimit    uint64
	BlockNumber *big.Int
	Time        *big.Int
	Difficulty  *big.Int
	BaseFee     *big.Int
}

type Context interface {
	Set(string, string) error
	String(string) string
	Bool(string) bool
}

type Trie interface {
	GetKey([]byte) []byte
	GetAccount(address Address) (*StateAccount, error)
	Hash() Hash
	NodeIterator(startKey []byte) NodeIterator
	Prove(key []byte, fromLevel uint, proofDb KeyValueWriter) error
}

type StateAccount struct {
	Nonce    uint64
	Balance  *big.Int
	Root     Hash // merkle root of the storage trie
	CodeHash []byte
}

type NodeIterator interface {
	Next(bool) bool
	Error() error
	Hash() Hash
	Parent() Hash
	Path() []byte
	NodeBlob() []byte
	Leaf() bool
	LeafKey() []byte
	LeafBlob() []byte
	LeafProof() [][]byte
	AddResolver(NodeResolver)
}

type NodeResolver func(owner Hash, path []byte, hash Hash) []byte

type KeyValueWriter interface {
	Put(key []byte, value []byte) error
	Delete(key []byte) error
}
