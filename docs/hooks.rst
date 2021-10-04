.. _hooks:

============
Plugin Hooks
============

The key to understanding a plugin is understanding its corresponding hook. Broadly, plugin hooks can be thought of as functions which deliver a function signiture to the **plugin loader**. The signature matches that of the function(s) called by the plugin and describes the data that is being captured and how the plugin will deliver said data upon invocation. 

The public plugin hook function should follow the naming convention ``Plugin$HookName``. The first argument should be a `` core.PluginLoader``, followed by any arguments required by the functions to be provided by any plugins implementing this hook.

**Public and Private Hook Functions**

Each hook provides both a Public and private version. The private plugin hook function should bear the same name as the public plugin hook function, but with a lower case first letter. The signature should match the public plugin hook function, except that the first argument referencing the PluginLoader should be removed. It should invoke the public plugin hook function on plugins.DefaultPluginLoader. It should always verify that the DefaultPluginLoader is non-nil, log warning and return if the DefaultPluginLoader has not been initialized.

**Invocation**

the private plugin hook function should be invoked, with the appropriate arguments, in a single line within the Geth codebase inorder to minimize unexpected conflicts merging upstream Geth into plugeth.

Selected Hooks
--------------


`StateUpdate`_ 
************

`Invocation`_

The state update plugin provides a snapshot of the state subsystem in the form of a a stateUpdate object. The stateUpdate object contains all information transformed by a transaction but not the transaction itself. 

Invoked for each new block, StateUpdate provides the changes to the blockchain state. root corresponds to the state root of the new block. parentRoot corresponds to the state root of the parent block. destructs serves as a set of accounts that self-destructed in this block. accounts maps the hash of each account address to the SlimRLP encoding of the account data. storage maps the hash of each account to a map of that account's stored data.

.. warning:: StateUpdate is only called if Geth is running with 
             ``-snapshots=true``. This is the default behavior for Geth, but if you are explicitly running with ``--snapshot=false`` this function will not be invoked. 
			

`AppendAncient`_
*************

Invoked when the freezer moves a block from LevelDB to the ancients database. ``number`` is the number of the block. ``hash`` is the 32 byte hash of the block as a raw ``[]byte``. ``header``, ``body``, and ``receipts`` are the RLP encoded versions of their respective block elements. ``td`` is the byte encoded total difficulty of the block.

`NewHead`_
**********

Invoked when a new block becomes the canonical latest block. Note that if several blocks are processed in a group (such as during a reorg) this may not be called for each block. You should track the prior latest head if you need to process intermediate blocks.

`GetRPCCalls`_
************

Invoked when the RPC handler registers a method call. returns the call ``id``, method ``name``, and any ``params`` that may have been passed in.  



**I am struggling all of assudden with if this is really this best course of action. Giving all of this reference, seems to beg the -utils vs plugeth conversation and I am not sure that is a smart thing to get into.**






.. _StateUpdate: https://github.com/openrelayxyz/plugeth/blob/develop/core/state/plugin_hooks.go
.. _Invocation: https://github.com/openrelayxyz/plugeth/blob/develop/core/state/statedb.go#L955
.. _AppendAncient: https://github.com/openrelayxyz/plugeth/blob/develop/core/rawdb/plugin_hooks.go
.. _GetRPCCalls: https://github.com/openrelayxyz/plugeth/blob/develop/rpc/plugin_hooks.go
.. _NewHead: https://github.com/openrelayxyz/plugeth/blob/develop/core/plugin_hooks.go#L108
