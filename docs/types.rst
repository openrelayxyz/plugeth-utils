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

A subcommand redifines the total behavior of Geth and could stand on its own.

Tracers
-------

Tracers are used to collect information from transaction execution, through Geth's ``debug_traceCall``, ``debug_traceTransaction``, and ``debug_traceBlock``. While standard Geth allows you to specify custom tracers in JavaScript, tracers written as plugins can be made much more performant.

While normal tracers run in the context of an RPC call, Live Tracers run on transactions in the course of Geth's block validation process. The methods for a tracer and live tracer are the same, but while tracers expose information through a RPC calls, live tracers simply provide an opportunity to aggregate information and have no inherent method for providing it to a consumer.



Subscriptions
-------------

A subscription must take a context.context as an argument and return a channel and an error. Subscriptions require a stable connection and return a stream of information. Subscriptions require a websocket connection and pass a json argument such as: ``{"jsonrpc":"2.0", "id": 0, "method": "namespace_subscribe", "params": ["subscriptionName", $args...]}``

.. NOTE:: Plugins are not limited to a singular functionality and can be customized to operate as hybrids of the above archtypes.

**todo: this page needs a lot of work**
