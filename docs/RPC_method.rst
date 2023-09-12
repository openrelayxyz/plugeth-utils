.. _RPC_method:

===========
RPC Methods
===========

GetAPIs
*******

For **RPC Methods** a Get APIs method is required in the body of the plugin in order to make the plugin available. The bulk of the implementation will be in the MyService struct. MyService should be a struct which includes two public functions. 

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

RPC Method
**********
(**accurate heading?**)

For RPC calls, a function should have a ``context.Context`` object as the first argument, followed by an arbitrary number of JSON marshallable arguments, and return either a single JSON marshal object, or a JSON marshallable object and an error. The RPC framework will take care of decoding inputs to this function and encoding outputs, and if the error is non-nil it will serve an error response.

A simple implimentation would look like so: 

**eventual link to documentation for hello or some other rpc plugin**

.. code-block:: Go

   func (h *MyService) HelloWorld(ctx context.Context) string {
     return "Hello World"
   }

.. Note:: For plugins such as RPC Methods whcih impliment a 
       GetAPIs function, an **Initialize Node** function may not be necesary as the ``core.Node`` and ``core.Backend`` will be made available with GetAPIs.

Access
******

As with pre-built plugins, a``.so`` will need to be built from``main.go`` and moved into ``~/.ethereum/plugins``. Geth will need to be started with with a ``http.api=mynamespace`` flag. Additionally you will need to include a ``--http`` flag in order to access the standard json rpc methods.

The plugin can now be accessed with an rpc call to ``mynamespace_helloWorld``.
