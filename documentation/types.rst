.. _types:

======================
Basic Types of Plugins
======================

RPC Methods
-----------

In general these plugins provide new json rpc methods. They will requirre an initialize function that takes a context, loader, and logger as arguments. They will also need a GetAPIs function that takes a node and backend as arguments and returns an API. 

.. NOTE:: In order to be made available a flag: ``http.api=<the name of your service>`` will need to be appended to the command line upon starting Geth. 

Subcommand
------------

A subcommand redifines the total behavior of Geth and could stand on its own. I contrast with the other plugin types which, in general, are meant to capture and manipulate information, a subcommand is meant to change to overall behavior of Geth. It may do this in order to capture information but the primary fuctionality is a modulation of geth behaviour. 

Tracers
-------

**Tracers vs LiveTracers. 

Tracers rely on historic data whereas LiveTracers run concurent with Geth's verification system. LiveTracers are more generic in that they cannot control the way in which they recieve information. 


Subscriptions
-------------

A subscription must take a context.context as an argument and return a channel and an error. Subscriptions require a stable connection and return contant information. Subscriptions require a websocket connection and pass a json argument such as: ``{"jsonrpc":"2.0", "id": 0, "method": "namespace_subscribe", "params": ["subscriptionName", $args...]}``

.. NOTE:: Plugins are not limited to a singular functionality and can be customized to operate as hybrids of the above archtypes. 

**todo: this page needs a lot of work**