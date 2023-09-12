.. _api:

===
API
===

Anatomy of a Plugin
===================

Plugins for Plugeth use Golang's `Native Plugin System`_. Plugin modules must export variables using specific names and types. These will be processed by the plugin loader, and invoked at certain points during Geth's operations.

Flags
-----

* **Name:** Flags
* **Type:** `flag.FlagSet`_
* **Behavior:** This FlagSet will be parsed and your plugin will be able to access the resulting flags. Flags will be passed to Geth from the command line and are intended to configure the behavior of the plugin. Passed flags must follow ``--`` to be parsed by this FlagSet, which is necessary to avoid Geth failing due to unexpected flags.

Subcommands
-----------

* **Name:** Subcommands
* **Type:** map[string]func(ctx `*cli.Context`_, args []string) error
* **Behavior:** If Geth is invoked with ``./geth YOUR_COMMAND``, the plugin loader will look for ``YOUR_COMMAND`` within this map, and invoke the corresponding function. This can be useful for certain behaviors like manipulating Geth's database without having to build a separate binary.

Initialize
----------

* **Name:** Initialize
* **Type:** func(*cli.Context, core.PluginLoader, core.logs )
* **Behavior:** Called as soon as the plugin is loaded, with the cli context and a reference to the plugin loader. This is your plugin's opportunity to initialize required variables as needed. Note that using the context object you can check arguments, and optionally can manipulate arguments if needed for your plugin. 

InitializeNode
--------------

* **Name:** InitializeNode
* **Type:** func(core.Node, core.Backend)
* **Behavior:** This is called as soon as the Geth node is initialized. The core.Node object represents the running node with p2p and RPC capabilities, while the Backend gives you access to a wide array of data you may need to access.

.. note:: If a particular plugin requires access to the node.Node object it can be obtained using the restricted package located in `PluGeth-Utils`_.

GetAPIs
-------

* **Name:** GetAPIs
* **Type:** func(core.Node, core.Backend) []rpc.API
* **Behavior:** This allows you to register new RPC methods to run within Geth.

The GetAPIs function itself will generally be fairly brief, and will looks something like this:

.. code-block:: go

	``func GetAPIs(stack *node.Node, backend core.Backend) []core.API {
        return []rpc.API{
         {
           Namespace: "mynamespace",
           Version:	 "1.0",
           Service:	 &MyService{backend},
           Public:		true,
         },
        }
      }``

The bulk of the implementation will be in the ``MyService`` struct. MyService should be a struct with public functions. These functions can have two different types of signatures:

* RPC Calls: For straight RPC calls, a function should have a ``context.Context`` object as the first argument, followed by an arbitrary number of JSON marshallable arguments, and return either a single JSON marshal object, or a JSON marshallable object and an error. The RPC framework will take care of decoding inputs to this function and encoding outputs, and if the error is non-nil it will serve an error response.

* Subscriptions: For subscriptions (supported on IPC and websockets), a function should have a ``context.Context`` object as the first argument followed by an arbitrary number of JSON marshallable arguments, and should return an ``*rpc.Subscription`` object. The subscription object can be created with ``rpcSub := notifier.CreateSubscription()``, and JSON marshallable data can be sent to the subscriber with ``notifier.Notify(rpcSub.ID, b)``.

A very simple MyService might look like:

.. code-block:: go

	``type MyService struct{}

	  func (h MyService) HelloWorld(ctx context.Context) string {
	    return "Hello World"
	  }``

And the client could access this with an rpc call to 
``mynamespace_helloworld``

Injected APIs
=============

In addition to hooks that get invoked by Geth, several objects are injected that give you access to additional information.

Backend Object
--------------

The ``core.Backend`` object is injected by the ``InitializeNode()`` and ``GetAPI()`` functions. It offers the following functions:

Downloader
**********
``Downloader() Downloader``

Returns a Downloader objects, which can provide Syncing status

SuggestGasTipCap
****************
``SuggestGasTipCap(ctx context.Context) (*big.Int, error)``

Suggests a Gas tip for the current block.

ExtRPCEnabled
*************
``ExtRPCEnabled() bool``

Returns whether RPC external RPC calls are enabled.

RPCGasCap
*********
``RPCGasCap() uint64``

Returns the maximum Gas available to RPC Calls.

RPCTxFeeCap
***********
``RPCTxFeeCap() float64``

Returns the maximum transaction fee for a transaction submitted via RPC.

UnprotectedAllowed
******************
``UnprotectedAllowed() bool``

Returns whether or not unprotected transactions can be transmitted through this
node via RPC.

SetHead
*******
``SetHead(number uint64)``

Resets the head to the specified block number.

HeaderByNumber
**************
``HeaderByNumber(ctx context.Context, number int64) ([]byte, error)``

Returns an RLP encoded block header for the specified block number.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Header` object.

HeaderByHash
************
``HeaderByHash(ctx context.Context, hash Hash) ([]byte, error)``

Returns an RLP encoded block header for the specified block hash.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Header` object.

CurrentHeader
*************
``CurrentHeader() []byte``

Returns an RLP encoded block header for the current block.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Header` object.

CurrentBlock
************
``CurrentBlock() []byte``

Returns an RLP encoded full block for the current block.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Block` object.


BlockByNumber
*************
``BlockByNumber(ctx context.Context, number int64) ([]byte, error)``


Returns an RLP encoded full block for the specified block number.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Block` object.

BlockByHash
***********
``BlockByHash(ctx context.Context, hash Hash) ([]byte, error)``

Returns an RLP encoded full block for the specified block hash.

The RLP encoded response can be decoded into a `plugeth-utils/restricted/types.Block` object.

GetReceipts
***********
``GetReceipts(ctx context.Context, hash Hash) ([]byte, error)``

Returns an JSON encoded list of receipts for the specified block hash.

The JSON encoded response can be decoded into a `plugeth-utils/restricted/types.Receipts` object.


GetTd
*****
``GetTd(ctx context.Context, hash Hash) *big.Int``

Returns the total difficulty for the specified block hash.

SubscribeChainEvent
*******************
``SubscribeChainEvent(ch chan<- ChainEvent) Subscription``

Subscribes the provided channel to new chain events.

SubscribeChainHeadEvent
***********************
``SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) Subscription``

Subscribes the provided channel to new chain head events.

SubscribeChainSideEvent
***********************
``SubscribeChainSideEvent(ch chan<- ChainSideEvent) Subscription``

Subscribes the provided channel to new chain side events.

SendTx
******
``SendTx(ctx context.Context, signedTx []byte) error``

Sends an RLP encoded, signed transaction to the network.

GetTransaction
**************
``GetTransaction(ctx context.Context, txHash Hash) ([]byte, Hash, uint64, uint64, error)``

Returns an RLP encoded transaction at the specified hash, along with the hash and number of the included block, and the transaction's position within that block.

GetPoolTransactions
^^^^^^^^^^^^^^^^^^^
``GetPoolTransactions() ([][]byte, error)``

Returns a list of RLP encoded transactions found in the mempool

GetPoolTransaction
******************
``GetPoolTransaction(txHash Hash) []byte``

Returns the RLP encoded transaction from the mempool at the specified hash.

GetPoolNonce
************
``GetPoolNonce(ctx context.Context, addr Address) (uint64, error)``

Returns the nonce of the last transaction for a given address, including
transactions found in the mempool.

Stats
*****
``Stats() (pending int, queued int)``

Returns the number of pending and queued transactions in the mempool.

TxPoolContent
*************
``TxPoolContent() (map[Address][][]byte, map[Address][][]byte)``

Returns a map of addresses to the list of RLP encoded transactions pending in
the mempool, and queued in the mempool.

SubscribeNewTxsEvent
********************
``SubscribeNewTxsEvent(chan<- NewTxsEvent) Subscription``

Subscribe to a feed of new transactions added to the mempool.

GetLogs
*******
``GetLogs(ctx context.Context, blockHash Hash) ([][]byte, error)``

Returns a list of RLP encoded logs found in the specified block.

SubscribeLogsEvent
******************
``SubscribeLogsEvent(ch chan<- [][]byte) Subscription``

Subscribe to logs included in a confirmed block.

SubscribePendingLogsEvent
*************************
``SubscribePendingLogsEvent(ch chan<- [][]byte) Subscription``

Subscribe to logs from pending transactions.

SubscribeRemovedLogsEvent
*************************
``SubscribeRemovedLogsEvent(ch chan<- []byte) Subscription``

Subscribe to logs removed from the canonical chain in reorged blocks.


Node Object
-----------

The ``core.Node`` object is injected by the ``InitializeNode()`` and ``GetAPI()`` functions. It offers the following functions:

Server
******
``Server() Server``

The Server object provides access to ``server.PeerCount()``, the number of peers connected to the node.

DataDir
*******
``DataDir() string``

Returns the Ethereuem datadir.

InstanceDir
***********
``InstanceDir() string``

Returns the instancedir used by the protocol stack.

IPCEndpoint
***********
``IPCEndpoint() string``

The path of the IPC Endpoint for this node.

HTTPEndpoint
************
``HTTPEndpoint() string``

The url of the HTTP Endpoint for this node.

WSEndpoint
**********
``WSEndpoint() string``

The url of the websockets Endpoint for this node.


ResolvePath
***********
``ResolvePath(x string) string``

Resolves a path within the DataDir.


.. _*cli.Context: https://pkg.go.dev/github.com/urfave/cli#Context
.. _flag.FlagSet: https://pkg.go.dev/flag#FlagSet
.. _Native Plugin System: https://pkg.go.dev/plugin

Logger
------

The Logger object is injected by the ``Initialize()`` function. It implements
logging based on the interfaces of `Log15 <https://github.com/inconshreveable/log15>`_.




.. _PluGeth-Utils: https://github.com/openrelayxyz/plugeth-utils
.. _*cli.Context: https://pkg.go.dev/github.com/urfave/cli#Context
.. _flag.FlagSet: https://pkg.go.dev/flag#FlagSet
.. _Native Plugin System: https://pkg.go.dev/plugin
