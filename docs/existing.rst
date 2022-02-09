.. _existing:

================
Existing Pluings
================

The isSynced plugin was designed as an extention of the ``eth_syncing`` available on statdard Geth. ``plugeth_isSynced`` was desinged to return a status object such that a status report could be given as to the current state of the node as opposed to  ``eth_syncing`` which returns the status object only if the node is actively syncing and a simple false if frozen or fully synced.    


Usage
======
As with all ``rpc`` methods, isSynced is available by ``curl`` or the `javascript console`_. 

From the command line: 

``{"method": "plugeth_isSynced", "params": []}``



``Lookup(name string, validate func(interface{}) bool) []interface{}``

Returns a list of values from plugins identified by ``name``, which match the
provided ``validate`` predicate. For example:


.. code-block:: go

    pl.Lookup("Version", func(item interface{}) bool {
      _, ok := item.(int)
      return ok
    })

Would return a list of ``int`` objects named ``Version`` in any loaded plugins.
This can enable Plugins to interact with each other, accessing values and
functions implemented in other plugins.

GetFeed
=======
``GetFeed() Feed``

.. _javascript console: https://geth.ethereum.org/docs/interface/javascript-console