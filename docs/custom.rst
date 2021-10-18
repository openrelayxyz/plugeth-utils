.. _custom:


========================
Building a Custom Plugin
========================

.. toctree::
    :hidden:

    RPC_method
    subscription
    tracer


Before setting out to build a plugin it will be helpful to be familiar with the :ref:`types`. deifferent plugins will require different implimentation.

Basic Implementation
====================

In general, no matter which type of plugin you intend to build, all will share some common aspects. 

Package
-------

Any plugin will need its own package located in the Plugeth-Plugins packages directory. The package will need to include a main.go from which the .so file will be built. The package and main file should share the same name and the name should be a word that describes the basic functionality of the plugin. 

Initialize
----------

Most plugins will need to be initialized with an Initialize function. The initialize function will need to be passed at least three arguments: a cli.Context, core.PluginLoader, and a core.Logger.  

And so, all plugins could have an intial template that looks something like this: 

.. code-block:: Go

   package main

   import (
	   "github.com/openrelayxyz/plugeth-utils/core"
	   "gopkg.in/urfave/cli.v1"
   )

   var log core.Logger

   func Initialize(ctx *cli.Context, loader core.PluginLoader, logger core.Logger) {
	   log = logger
	   log.Info("loaded New Custom Plugin")
   }

InitializeNode
--------------

Many plugins will make use of the InitializeNode function. Implimentation will look like so:

.. code-block:: Go

   func InitializeNode(stack core.Node, b core.Backend) {
           backend = b
           log.Info("Initialized node and backend")
   }


This is called as soon as the Geth node is initialized. The core.Node object represents the running node with p2p and RPC capabilities, while the Backend gives you access to blocks and other data you may need to access.    

Specialization
==============

From this point implimentation becomes more specialized to the particular plugin type. Continue from here for specific instructions for the following plugins:

* :ref:`RPC_method`
* :ref:`subscription`
* :ref:`tracer`








.. _blockupdates: https://github.com/openrelayxyz/plugeth-plugins/blob/master/packages/blockupdates/main.go
.. _hello: https://github.com/openrelayxyz/plugeth-plugins/blob/master/packages/hello/main.go
