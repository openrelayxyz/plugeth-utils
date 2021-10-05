.. _build:

================
Build and Deploy
================

.. contents:: :local:

If you are ready to start building your own plugins go ahead and start here. 

Setting up the environment
**************************

.. NOTE:: Plugeth is built on a fork of `Geth`_ and as such requires familiarity with `Go`_ and a funtional `environment`_ in which to build Go projects. Thankfully for everyone Go provides a compact and useful `tutorial`_ as well as a `space for practice`_. 

PluGeth is an application built in three seperate repositories. 

* `PluGeth`_
* `PluGeth-Utils`_
* `PluGeth-Plugins`_

For our purposed here you will only need to clone Plugeth and Plugeth-Plugins. Once you have them cloned you are ready to begin. First we need to build geth though the PluGeth project. Navigate to ``plugeth/cmd/geth`` and run:

.. code-block:: shell

   $ go get

This will download all dependencies needed for the project. This process will take a moment or two the first time through. Next run: 

.. code-block:: shell

   $ go build

Once this is complete you should see a ``.ethereum`` folder in your home directory. 

At this point you are ready to start downloading local ethereum nodes. In order to do so, from ``plugeth/cmd/geth`` run:

.. code-block:: shell

   $ ./geth

.. NOTE:: ``./geth`` is the primary command to build a *Mainnet* node. Building a mainnet node requires at least 8 GB RAM, 2 CPUs, and 350 GB of SSD disks. However, dozens of available flags will change the behavior of whichever network you choose to connect to. ``--help`` is your friend. 


Build a plugin
**************

For the sake of this tutorial we will be building the Hello plugin. Navigate to ``plugethPlugins/packages/hello``. Inside you will see a ``main.go`` file. From this location run:

.. code-block:: shell

   $ go build -buildmode=plugin

This will compile the plugin and produce a ``hello.so`` file. Move ``hello.so`` into ``~/.ethereum/plugins`` . In order to use this plugin geth will need to be started with a ``http.api=mymamespace`` flag. Additionally you will need to include a ``--http`` flag in order to access the standard json rpc methods. 

Once geth has started you should see that the first ``INFO`` log reads: ``initialized hello`` . A new json rpc method, called hello, has been been appended to the list of available json rpc methods. In order to access this method you will need to ``curl`` into the network with this command:

.. code-block:: shell

   $ curl 127.0.0.1:8545 -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"mynamespace_hello","params":[],"id":0}'

You should see that the network has responded with:

.. code-block:: shell

   ``{"jsonrpc":"2.0","id":0,"result":"Hello world"}``

Congradulations. You have just built and run your first Plugeth plugin. From here you can follow the steps above to build any of the plugins you choose. 

.. NOTE:: Each plugin will vary in terms of the requirements to deploy. Refer to the documentation of the plugin itself in order to assure 
          that you know how to use it. 



.. _space for practice: https://tour.golang.org/welcome/1 
.. _tutorial: https://tour.golang.org/welcome/1 
.. _environment: https://golang.org/doc/code
.. _Go: https://golang.org/doc/
.. _Geth: https://geth.ethereum.org/
.. _PluGeth: https://github.com/openrelayxyz/plugeth
.. _PluGeth-Utils: https://github.com/openrelayxyz/plugeth-utils
.. _PluGeth-Plugins: https://github.com/openrelayxyz/plugeth-plugins