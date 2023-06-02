.. _tracer:

======
Tracer
======

In addition to the initial template containing an intialize function, plugins providing **Tracers** will require three additional elements. 

.. Warning:: Caution: Modifying of the values passed into tracer 
             functions can alter the results of the EVM execution in unpredictable ways. Additionally, some objects may be reused across calls, so data you wish to capture should be copied rather than retained be reference. 

MyService Struct
****************

First an empty MyService Struct.

.. code-block:: Go

   type MyService struct {
   }

Map
***
   
Next, a map of tracers to functions returning a ``core.TracerResult`` which will be implemented like so:

.. code-block:: Go
   
   var Tracers = map[string]func(core.StateDB) core.TracerResult{
       "myTracer": func(core.StateDB) core.TracerResult {
           return &MyBasicTracerService{}
       },
   }

TracerResult Functions
**********************

Finally a series of functions which points to the MyService struct and coresponds to the interface which geth anticipates.

.. code-block:: Go

   func (b *MyBasicTracerService) CaptureStart(from core.Address, to core.Address, create bool, input []byte, gas uint64, value *big.Int) {
   }
   func (b *MyBasicTracerService) CaptureState(pc uint64, op core.OpCode, gas, cost uint64, scope core.ScopeContext, rData []byte, depth int, err error) {
   }
   func (b *MyBasicTracerService) CaptureFault(pc uint64, op core.OpCode, gas, cost uint64, scope core.ScopeContext, depth int, err error) {
   }
   func (b *MyBasicTracerService) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) {
   } 
   func (b *MyBasicTracerService) CaptureEnter(typ core.OpCode, from core.Address, to core.Address, input []byte, gas uint64, value *big.Int) {
   }
   func (b *MyBasicTracerService) CaptureExit(output []byte, gasUsed uint64, err error) {
   }
   func (b *MyBasicTracerService) Result() (interface{}, error) { return "hello world", nil }

Access
******
As with pre-built plugins, a ``.so`` will need to be built from ``main.go`` and moved into ``~/.ethereum/plugins``. Geth will need to be started with with a ``--http.api+debug`` flag. 

From a terminal pass the following argument to the api:

.. code-block:: shell

   curl 127.0.0.1:8545 -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"debug_traceCall","params":[{"to":"0x32Be343B94f860124dC4fEe278FDCBD38C102D88"},"latest",{"tracer":"myTracer"}],"id":0}'
   
.. Note:: The address used above is a test adress and will need to
          be replaced by whatever address you wish to access. Also ``traceCall`` is one of several methods available for use. 

If using the template above, the call should return:

.. code-block:: shell

   {"jsonrpc":"2.0","id":0,"result":"hello world"}


