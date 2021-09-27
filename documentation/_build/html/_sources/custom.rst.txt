.. _custom:

========================
Building a Custom Plugin
========================


Before setting out to build a plugin it will be helpful to be familiar with the basic :ref:`types` of plugins. Depending on what you intend to do certain aspects of implimentation will be neccisary.

In general, no matter which type of plugin you intend to build, below are common aspects which will be shared by all plugins.


Most Basic implimentation
=========================


A plugin will need its own package located in the Plugeth-Utils packages directory. The package will need to include a main.go from which the .so file will be built. The package and main file should share the same name and the name should be a word that describes the basic functionality of the plugin. 

All plugins will need to be initialized with an **initialize function**. The initialize function will need to be passed at least three arguments: a cli.Context, core.PluginLoader, and a core.Logger.  

And so, all plugins will have an intial template that looks something like this: 

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

Specialization
==============

**Hooks**

Plugeth provides several hooks with which data of various kinds can be captured and manipulated. Once a plugin has been initalized it will be up to the hooks utilized to determine the behavior of the plugin.    


