.. _project:

==============
Project Design
==============

Design Goals

The upstream Geth client exists primarily to serve as a client for the Ethereum mainnet, though it also supports a number of popular testnets. Supporting the Ethereum mainnet is a big enough challenge in its own right that the Geth team generally avoids changes to support other networks, or to provide features only a small handful of users would be interested in.

The result is that many projects have forked Geth. Some implement their own consensus protocols or alter the behavior of the EVM to support other networks. Others are designed to extract information from the Ethereum mainnet in ways the standard Geth client does not support.

Creating numerous different forks to fill a variety of different needs comes with a number of drawbacks. Forks tend to drift apart from each other. Many networks that forked from Geth long ago have stopped merging updates from Geth; this makes some sense, given that those networks have moved in different directions than Geth and merging upstream changes while properly maintaining consensus rules of an existing network could prove quite challenging. But not merging changes from upstream can mean that security updates are easily missed, especially when the upstream team `obscures security updates as optimizations`_ as a matter of process.

PluGeth aims to provide a single Geth fork that developers can choose to extend rather than forking the Geth project. Out of the box, PluGeth behaves exactly like upstream Geth, but by installing plugins written in Golang, developers can extend its functionality in a wide variety of ways.

Three Repositories
------------------

PluGeth is an application built in three repositories:

`PluGeth`_
**********

The largest of the three Repositories, PluGeth is a fork of Geth which has been modified to enable a plugin architecture. The Plugin loader, wrappers, and hooks all reside in this repository.

`PluGeth-Utils`_
***************

Utils are small packages used to develop PluGeth plugins without Geth dependencies. For a more detailed analysis of the reasons see **here**

`PluGeth-Plugins`_
*****************

Plugins are packages which contain premade plugins as well as a location provided for storing new custom plugins.

Dependency Scheme
-----------------

PluGeth is separated into three packages in order to minimize dependency conflicts. Golang plugins cannot include different versions of the same packages as the program loading the plugin. If plugins had to import packages from PluGeth itself, a plugin build could only be loaded by that same version of PluGeth. By separating out the PluGeth-utils package, both PluGeth and the plugins must rely on the same version of PluGeth-utils, but plugins can be compatible with any version of PluGeth compiled with the same version of PluGeth-utils.

PluGeth builds will follow the naming convention:

.. code-block:: shell

   geth-$PLUGETH_UTILS_VERSION-$GETH_VERSION-$RELEASE

For example:

.. code-block:: shell

   geth-0.1.0-1.10.8-0

Tells us that:

* PluGeth-utils version is 0.1.0
* Geth version is 1.10.8
* This is the first release with that combination of dependencies.

Plugin builds will follow the naming convention:

.. code-block:: shell

   $PLUGIN_NAME-$PLUGETH_UTILS_VERSION-$PLUGIN_VERSION

For example:

.. code-block:: shell

   blockupdates-0.1.0-1.0.2

Tells us that:

* The plugin is "blockupdates"
* The PluGeth-utils version is 0.1.0
* The plugin version is 1.0.2

When a Geth update comes out, you can expect a release of `geth-0.1.0-1.10.9-0`, which will be compatible with the same set of plugins.

When PluGeth upgrades are necessary, plugins will need to be recompiled. Whenever possible, we will try to avoid forcing plugins to be recompiled for an immediate Geth upgrade. For example, if we have geth-0.1.0-1.10.8, and upgrade PluGeth-utils, we will have a geth-0.1.1-1.10.8, followed by a geth-0.1.1-1.10.9. This will give users time to upgrade plugins from PluGeth-utils 0.1.0 to 0.1.1 while staying on Geth 1.10.8, and when it is time to upgrade to Geth 1.10.9 they can continue using the plugins they were using with geth 1.10.8. Depending on upgrades to Geth, it may not always be possible to maintain compatibility with existing PluGeth versions, which will be noted in release notes.


.. _obscures security updates as optimizations: https://blog.openrelay.xyz/vulnerability-lifecycle-framework-geth/
.. _PluGeth: https://github.com/openrelayxyz/plugeth
.. _PluGeth-Utils: https://github.com/openrelayxyz/plugeth-utils
.. _PluGeth-Plugins: https://github.com/openrelayxyz/plugeth-plugin
