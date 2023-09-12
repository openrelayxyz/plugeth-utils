.. _hooks:

=====================
Selected Plugin Hooks
=====================

Plugin Hooks
************

Plugeth provides several hooks from which the plugin can capture data from Geth. Additionally in the case of **subcommands** the provided hooks are designed to change the behavior of Geth.

Hooks are called from functions within the plugin. For example, if we wanted to bring in data from the StateUpdate hook. We would impliment it like so:
(from `blockupdates`_)

.. code-block:: Go

   func StateUpdate(blockRoot core.Hash, parentRoot core.Hash, destructs map[core.Hash]struct{}, accounts map[core.Hash][]byte, storage map[core.Hash]map[core.Hash][]byte, codeUpdates map[core.Hash][]byte) {
         su := &stateUpdate{
                 Destructs: destructs,
                 Accounts: accounts,
                 Storage: storage,
                 Code: codeUpdates,
         }
         cache.Add(blockRoot, su)
         data, _ := rlp.EncodeToBytes(su)
         backend.ChainDb().Put(append([]byte("su"), blockRoot.Bytes()...), data)
   }

Many hooks can be deployed in one plugin as is the case with the **BlockUpdater** plugin.

.. contents:: :local:



StateUpdate
***********

**Function Signature**:``func(root common.Hash, parentRoot common.Hash, destructs map[common.Hash]struct{}, accounts map[common.Hash][]byte, storage map[common.Hash]map[common.Hash][]byte)``

The state update plugin provides a snapshot of the state subsystem in the form of a a stateUpdate object. The stateUpdate object contains all information transformed by a transaction but not the transaction itself.

Invoked for each new block, StateUpdate provides the changes to the blockchain state. root corresponds to the state root of the new block. parentRoot corresponds to the state root of the parent block. destructs serves as a set of accounts that self-destructed in this block. accounts maps the hash of each account address to the SlimRLP encoding of the account data. storage maps the hash of each account to a map of that account's stored data.

.. warning:: StateUpdate is only called if Geth is running with
             ``-snapshots=true``. This is the default behavior for Geth, but if you are explicitly running with ``--snapshot=false`` this function will not be invoked.


AppendAncient
*************

**Function Signature**:``func(number uint64, hash, header, body, receipts, td []byte)``

Invoked when the freezer moves a block from LevelDB to the ancients database. ``number`` is the number of the block. ``hash`` is the 32 byte hash of the block as a raw ``[]byte``. ``header``, ``body``, and ``receipts`` are the RLP encoded versions of their respective block elements. ``td`` is the byte encoded total difficulty of the block.

GetRPCCalls
***********

**Function Signature**:``func(string, string, string)``

Invoked when the RPC handler registers a method call. Returns the call ``id``, method ``name``, and any ``params`` that may have been passed in.

.. todo:: missing a couple of hooks

PreProcessBlock
***************

**Function Signature**:``func(*types.Block)``

Invoked before the transactions of a block are processed. Returns a block object.

PreProcessTransaction
*********************

**Function Signature**:``func(*types.Transaction, *types.Block, int)``

Invoked before each individual transaction of a block is processed. Returns a transaction, block, and index number.

BlockProcessingError
********************

**Function Signature**:``func(*types.Transaction, *types.Block, error)``

Invoked if an error occurs while processing a transaction. This only applies to errors that would unvalidate the block were this transaction is included not errors such as reverts or opcode errors. Returns a transaction, block, and error.

NewHead
*******

**Function Signature**:``func(*types.Block, common.Hash, []*types.Log)``

Invoked when a new block becomes the canonical latest block. Returns a block, hash, and logs.

.. note:: If several blocks are processed in a group (such as
          during a reorg) this may not be called for each block. You should track the prior latest head if you need to process intermediate blocks.

NewSideBlock
************

**Function Signature**:``func(*types.Block, common.Hash, []*types.Log)``

Invoked when a block is side-chained. Returns a block, hash, and logs.

.. note:: Blocks passed to this method are non-canonical blocks.


Reorg
*****

**Function Signature**:``func(common *types.Block, oldChain, newChain types.Blocks)``

Invoked when a chain reorg occurs, that is; at least one block is removed and one block is added. (``oldChain`` is a list of removed blocks, ``newChain`` is a list of newliy added blocks, and ``common`` is the latest block that is an ancestor to both oldChain and newChain.) Returns a block, a list of old blocks, and a list of new blocks.





.. _blockupdates: https://github.com/openrelayxyz/plugeth-plugins/tree/master/packages/blockupdates
.. _StateUpdate: https://github.com/openrelayxyz/plugeth/blob/develop/core/state/plugin_hooks.go
.. _Invocation: https://github.com/openrelayxyz/plugeth/blob/develop/core/state/statedb.go#L955
.. _AppendAncient: https://github.com/openrelayxyz/plugeth/blob/develop/core/rawdb/plugin_hooks.go
.. _GetRPCCalls: https://github.com/openrelayxyz/plugeth/blob/develop/rpc/plugin_hooks.go
.. _NewHead: https://github.com/openrelayxyz/plugeth/blob/develop/core/plugin_hooks.go#L108
