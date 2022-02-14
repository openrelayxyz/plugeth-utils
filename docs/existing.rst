.. _existing:

================
Existing Pluings
================

getRPCCalls
===========

getRPCCalls is a subcommand written to print a log containing information about RPC methods upon execution. Namely the id, method name, and parameters into the method. 

Usage
-----

Once compiled the plugin will execute automatically as RPC methods are passed into the api. 

isSynced
========

The isSynced plugin was designed as an extention of the ``eth_syncing`` method available on standard Geth. ``plugeth_isSynced`` was desinged to return a status object such that a status report could be given as to the current state of the node as opposed to  ``eth_syncing`` which returns the status object only if the node is actively syncing and a simple false if frozen or fully synced.    


Usage
-----
As with all ``rpc`` methods, isSynced is available by ``curl`` or the `javascript console`_. 

From the command line using the ``curl`` command: 

``{"method": "plugeth_isSynced", "params": []}``

Which will return: 

.. code-block:: Json
   
    "activePeers": true,
    "currentBlock": "0x60e880",
    "healedBytecodeBytes": "0x0",
    "healedBytecodes": "0x0",
    "healedTrienodeBytes": "0x0",
    "healedTrienodes": "0x0",
    "healingBytecode": "0x0",
    "healingTrienodes": "0x0",
    "highestBlock": "0x60e880",
    "nodeIsSynced": true,
    "startingBlock": "0x0",
    "syncedAccountBytes": "0x0",
    "syncedAccounts": "0x0",
    "syncedBytecodeBytes": "0x0",
    "syncedBytecodes": "0x0",
    "syncedStorage": "0x0",
    "syncedStorageBytes": "0x0"




blockTracer
===========

Blocktracer is an subscription plugin written such that for each block mined, blockTracer will return a json payload reporting the type, from and to addresses, gas, gas used, input, output, and calls made for each transaction. The data will stream in real time as the block is mined. 

Usage
=====

As with any websocket an itial connection will need to be established. 


Here we are using wscat to connect to local host port 8556.

``wscat -c "http://127.0.0.1:8556"`` 

Once the connection has been made the method as well as blockTracer parameter will be passed in. 

``{"method":"plugeth_subscribe","params":["traceBlock"],"id":0}``

Which will return a streaming result similar to the one below. 

.. code-block:: Json

    "type":"CALL","from":"0x75d5e88adf8f3597c7c3e4a930544fb48089c779","to":"0x9ac40b4e6a0c60ca54a7fa2753d65448e6a71ecb","gas":"0x58cc2","gasUsed":"0x6007","input":"0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028d2f41e4c1dfca58114457fbe07632cabbfb9d900000000000000000000000000000000000000000000000000000000001db898fbdbdd5c","output":"0x0000000000000000000000000000000000000000000000000000000000000000","calls":[{"type":"DELEGATECALL","from":"0x9ac40b4e6a0c60ca54a7fa2753d65448e6a71ecb","to":"0xae9a8ae28d55325dff2af4ed5fe2335c1a39139b","gas":"0x56308","gasUsed":"0x4c07","input":"0x0000000000000000000000000000000000000000000000000000000000000000abbfb9d900000000000000000000000000000000000000000000000000000000001db8980000000000000000000000000000000000000000035298ac0ba8bb05fbdbdd5c","output":"0x0000000000000000000000000000000000000000000000000000000000000000"}]}]}]}}


.. _javascript console: https://geth.ethereum.org/docs/interface/javascript-console