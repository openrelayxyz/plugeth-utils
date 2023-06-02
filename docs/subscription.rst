.. _subscription:

============
Subscription
============

In addition to the initial template containing an intialize function, plugins providing **Subscriptions** will require two additional elements. 

GetAPIs
*******

A GetAPIs method is required in the body of the plugin in order to make the plugin available. The bulk of the implementation will be in the MyService struct. MyService should be a struct which includes two public functions. 

.. code-block:: Go

   type MyService struct {
       backend core.Backend
       stack   core.Node
   }

   func GetAPIs(stack core.Node, backend core.Backend) []core.API {
     return []core.API{
      {
         Namespace: "plugeth",
         Version:   "1.0",
         Service:   &MyService{backend, stack},
         Public:    true,
      },
    }
   }

Subscription Function
*********************

For subscriptions (supported on IPC and websockets), a function should take MyService as a reciever and a context.Context object as an argument and return a channel and an error. The following is a subscription function that implements a timer. 

.. code-block:: Go

   
   func (*myservice) Timer(ctx context.Context) (<-chan int64, error) {
           ticker := time.NewTicker(time.Second)
           ch := make(chan int64)
           go func() {
                   defer ticker.Stop()
                   for {
                           select {
                           case <-ctx.Done():
                                   close(ch)
                                   return
                           case t := <-ticker.C:
                                   ch <- t.UnixNano()
                           }
                   }
           }()
           return ch, nil
   }

.. warning:: Notice in the example above, the ``ctx.Done()`` or    
             Context.Done() method closes the channel. If this is not present the go routine will run for the life of the process. 

Access
******

.. Note:: Plugins providing subscriptions can be accessed via IPC 
          and websockets. In the below example we will be using `wscat`_ to connect a websocket to a local Geth node.

As with pre-built plugins, a ``.so`` will need to be built from ``main.go`` and moved into ``~/.ethereum/plugins``. Geth will need to be started with ``--ws --ws.api=mynamespace`` flags. Additionally you will need to include a ``--http`` flag in order to access the standard json rpc methods.

After starting Geth, from a seperate terminal run:

.. code-block:: shell

   wscat -c ws://127.0.0.1:8546

.. Note:: Websockets are available via port 8546

Once the connection has been established from the websocket cursor enter the following argument:

.. code-block:: shell

   {"jsonrpc":"2.0","method":"mynamespace_hello","params":[],"id":0}

   You should see that the network has responded with:

.. code-block:: shell

   ``{"jsonrpc":"2.0","id":0,"result":"Hello world"}``

.. _wscat: https://www.npmjs.com/package/wscat




