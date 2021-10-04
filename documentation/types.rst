.. _types:

======================
Basic Types of Plugins
======================

While PluGeth has been designed to be versatile and customizable, when learning the project it can be helpful to think of plugins as being of four different archetypes. 

.. contents:: :local:

RPC Methods
-----------

These plugins provide new json rpc methods to access several objects containing real time and historic data.


Subcommand
------------

A subcommand redifines the total behavior of Geth and could stand on its own. In contrast with the other plugin types which, in general, are meant to capture and manipulate information, a subcommand is meant to change to overall behavior of Geth. It may do this in order to capture information but the primary fuctionality is a modulation of geth behaviour. 

Tracers
-------

Tracers rely on historic data recompiled after execution to give insight into a transaction. 

**placeholder for eventual discusion of LiveTracers**


Subscriptions
-------------

Subscriptions provide real time notification of data from the EVM as it processes transactions. 

.. NOTE:: Plugins are not limited to a singular functionality and can be customized to operate as hybrids of the above. See `blockupdates`_ as an example. 

.. todo:: Austin: I don't love this page. The informations is too 
          shallow. 


.. _blockupdates: https://github.com/openrelayxyz/plugeth-plugins/tree/master/packages/blockupdates

