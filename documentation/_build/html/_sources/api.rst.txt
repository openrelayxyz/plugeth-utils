.. _api:

===
API
===

Plugins for Plugeth use Golang's `Native Plugin System`_. Plugin modules must export variables using specific names and types. These will be processed by the plugin loader, and invoked at certain points during Geth's operations.

Flags
-----

* **Name:** Flags
* **Type:** `flag.FlagSet`_
* **Behavior:** This FlagSet will be parsed and your plugin will be able to access the resulting flags. Flags will be passed to Geth from the command line and are intended to  of the plugin. Note that if any flags are provided, certain checks are disabled within Geth to avoid failing due to unexpected flags.

Subcommands
-----------

* **Name:** Subcommands
* **Type:** map[string]func(ctx `*cli.Context`_, args []string) error
* **Behavior:** If Geth is invoked with ``./geth YOUR_COMMAND``, the plugin loader will look for ``YOUR_COMMAND`` within this map, and invoke the corresponding function. This can be useful for certain behaviors like manipulating Geth's database without having to build a separate binary.

Initialize
----------

* **Name:** Initialize
* **Type:** func(*cli.Context, core.PluginLoader, core.logs )
* **Behavior:** Called as soon as the plugin is loaded, with the cli context and a reference to the plugin loader. This is your plugin's opportunity to initialize required variables as needed. Note that using the context object you can check arguments, and optionally can manipulate arguments if needed for your plugin. 

.. todo:: explain that plugin could provide node.Node with 
          restricted.backend

InitializeNode
--------------

* **Name:** InitializeNode
* **Type:** func(core.Node, core.Backend)
* **Behavior:** This is called as soon as the Geth node is initialized. The core.Node object represents the running node with p2p and RPC capabilities, while the Backend gives you access to a wide array of data you may need to access.

Tracers
-------

* **Name:** Tracers
* **Type:** map[string]TracerResult
* **Behavior:** When calling debug.traceX functions (such as ``debug_traceCall`` and ``debug_traceTransaction``) the tracer can be specified as a key to this map and the tracer used will be the TracerResult specified here. TracerResult objects must match the interface:

.. code-block:: go

   ``// CaptureStart is called at the start of each transaction
   CaptureStart(env core.EVM, from core.Address, to core.Address, create bool, input []byte, gas uint64, value *big.Int) {}
   // CaptureState is called for each opcode
   CaptureState(env core.EVM, pc uint64, op core.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {}
   // CaptureFault is called when an error occurs in the EVM
   CaptureFault(env core.EVM, pc uint64, op core.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {}
   // CaptureEnd is called at the end of each transaction
   CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) {}
   // GetResult should return a JSON serializable result object to respond to the trace call
   GetResult() (interface{}, error) {}``


.. warning:: Modifying the values passed into tracer functions can
             alter the 
             results of the EVM execution in unpredictable ways. Additonally, some objects may be reused acress calls, so data you wish to capture should be copied rather than retianed by reference. 

LiveTracer
----------

* **Name:** LiveTracers
* **Type:** core.Tracer
* **Behavior:** This tracer is used for tracing transactions as they are processed within blocks. Note that if a block does not validate, some transactions may be processed that don't end up in blocks, so be sure to check transactions against finalized blocks.

The interface for a vm.Tracer is similar to a TracerResult (above), but does not require a ``GetResult()`` function.

GetAPIs
-------

* **Name:** GetAPIs
* **Type:** func(core.Node, core.Backend) []rpc.API
* **Behavior:** This allows you to register new RPC methods to run within Geth.

The GetAPIs function itself will generally be fairly brief, and will looks something like this:

.. code-block:: go

	``func GetAPIs(stack *node.Node, backend core.Backend) []core.API {
        return []rpc.API{
         {
           Namespace: "mynamespace",
           Version:	 "1.0",
           Service:	 &MyService{backend},
           Public:		true,
         },
        }
      }``

The bulk of the implementation will be in the ``MyService`` struct. MyService should be a struct with public functions. These functions can have two different types of signatures:

* RPC Calls: For straight RPC calls, a function should have a ``context.Context`` object as the first argument, followed by an arbitrary number of JSON marshallable arguments, and return either a single JSON marshal object, or a JSON marshallable object and an error. The RPC framework will take care of decoding inputs to this function and encoding outputs, and if the error is non-nil it will serve an error response.

* Subscriptions: For subscriptions (supported on IPC and websockets), a function should have a ``context.Context`` object as the first argument followed by an arbitrary number of JSON marshallable arguments, and should return an ``*rpc.Subscription`` object. The subscription object can be created with ``rpcSub := notifier.CreateSubscription()``, and JSON marshallable data can be sent to the subscriber with ``notifier.Notify(rpcSub.ID, b)``.

A very simple MyService might look like:

.. code-block:: go

	``type MyService struct{}

	  func (h MyService) HelloWorld(ctx context.Context) string {
	    return "Hello World"
	  }``

And the client could access this with an rpc call to 
``mynamespace_helloworld``




.. _*cli.Context: https://pkg.go.dev/github.com/urfave/cli#Context
.. _flag.FlagSet: https://pkg.go.dev/flag#FlagSet
.. _Native Plugin System: https://pkg.go.dev/plugin