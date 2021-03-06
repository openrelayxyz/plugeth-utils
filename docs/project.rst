.. _project:

==============
Project Design
==============

Design Goals
============

Upstream Geth exists primarily to serve as a client for the Ethereum mainnet, though it also supports a number of popular testnets. The Geth team generally avoids changes to support other networks, or to provide features only a small handful of users would be interested in.

The result is that many projects have forked Geth. Some implement their own consensus protocols or alter the behavior of the EVM to support other networks. Others are designed to extract information from the Ethereum mainnet in ways the standard Geth client does not support.

Creating numerous different forks to fill a variety of different needs comes with a number of drawbacks. Forks tend to drift apart from each other. Many networks that forked from Geth long ago have stopped merging updates; this makes some sense, given that those networks have moved in different directions than Geth and merging upstream changes while properly maintaining consensus rules of an existing network could prove quite challenging. But not merging changes from upstream can mean that security updates are easily missed, especially when the upstream team `obscures security updates as optimizations`_ as a matter of process.

PluGeth aims to provide a single Geth fork that developers can choose to extend rather than forking the Geth project. Out of the box, PluGeth behaves exactly like upstream Geth, but by installing plugins written in Golang, developers can extend its functionality in a wide variety of ways.

Three Repositories
------------------

**PluGeth is an application built in three repositories:**

`PluGeth`_
----------

The largest of the three, PluGeth is a fork of Geth which has been modified to enable a plugin architecture. The Plugin loader, wrappers, and hooks all reside in this repository. 

`PluGeth-Utils`_
----------------

Utils are small packages used to develop PluGeth plugins without Geth dependencies. For a more detailed analysis of the reasons see :ref:`core_restricted`. Imports from Utils happen automatically and so most users need not clone a local version. 

`PluGeth-Plugins`_
------------------

The packages from which plugins are buile are stored here. This repository contains premade plugins as well as providing a location for storing new custom plugins. 

Version Control 
*****************

Before using Plugeth users are enocuraged to familiarize themselves with the :ref:`version control<version>` scheme of the project.  






.. _obscures security updates as optimizations: https://blog.openrelay.xyz/vulnerability-lifecycle-framework-geth/
.. _PluGeth: https://github.com/openrelayxyz/plugeth
.. _PluGeth-Utils: https://github.com/openrelayxyz/plugeth-utils
.. _PluGeth-Plugins: https://github.com/openrelayxyz/plugeth-plugins