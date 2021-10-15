.. _hook_writing:

==================
Hook Writing Guide
==================

If you're trying to interact with Geth in a way not already supported by
PluGeth, we're happy to accept pull requests adding new hooks so long as they
comply with certain standards. We strongly encourage you to :ref:`contact us <contact>`
first. We may have suggestions on how to do what you're trying to do without
adding new hooks, or easier ways to implement hooks to get the information you
need.

.. warning::

    Plugin hooks *must not* require plugins to import any packages from ``github.com/ethereum/go-ethereum``.
    Doing so means that plugins must be recompiled for each version of Geth.
    Many types have been re-implemented in ``github.com/openrelayxyz/plugeth-utils``.
    If you need a type for your hook not already provided by plugeth-utils, you
    may make a pull request to that project as well.

When extending the plugin API, a primary concern is leaving a minimal footprint
in the core Geth codebase to avoid future merge conflicts. To achieve this,
when we want to add a hook within some existing Geth code, we create a
plugin_hooks.go in the same package. For example, in the core/rawdb package we
have:


.. code-block:: Go

    // This file is part of the package we are adding hooks to
    package rawdb

    // Import whatever is necessary
    import (
      "github.com/ethereum/go-ethereum/plugins"
      "github.com/ethereum/go-ethereum/log"
    )


    // PluginAppendAncient is the public plugin hook function, available for testing
    func PluginAppendAncient(pl *plugins.PluginLoader, number uint64, hash, header, body, receipts, td []byte) {
      fnList := pl.Lookup("AppendAncient", func(item interface{}) bool {
        _, ok := item.(func(number uint64, hash, header, body, receipts, td []byte))
        return ok
      })
      for _, fni := range fnList {
        if fn, ok := fni.(func(number uint64, hash, header, body, receipts, td []byte)); ok {
          fn(number, hash, header, body, receipts, td)
        }
      }
    }

    // pluginAppendAncient is the private plugin hook function
    func pluginAppendAncient(number uint64, hash, header, body, receipts, td []byte) {
      if plugins.DefaultPluginLoader == nil {
    		log.Warn("Attempting AppendAncient, but default PluginLoader has not been initialized")
        return
      }
      PluginAppendAncient(plugins.DefaultPluginLoader, number, hash, header, body, receipts, td)
    }

The Public Plugin Hook Function
*******************************

The public plugin hook function should follow the naming convention
Plugin$HookName. The first argument should be a ``*plugins.PluginLoader``, followed
by any arguments required by the functions to be provided by nay plugins
implementing this hook.

The plugin hook function should use ``PluginLoader.Lookup("$HookName", func(item interface{}) bool``
to get a list of the plugin-provided functions to be invoked. The provided
function should verify that the provided function implements the expected
interface. After the first time a given hook is looked up through the plugin
loader, the PluginLoader will cache references to those hooks.

Given the function list provided by the plugin loader, the public plugin hook
function should iterate over the list, cast the elements to the appropriate
type, and call the function with the provided arguments.

Unless there is a clear justification to the contrary, the function should be
called in the current goroutine. Plugins may choose to spawn off a separate
goroutine as appropriate, but for the sake of thread safety we should generally
not assume that plugins will be implemented in a threadsafe manner. If a plugin
degrades the performance of Geth significantly, that will generally be obvious,
and plugin authors can take appropriate measures to improve performance. If a
plugin introduces thread safety issues, those can go unnoticed during testing.

The Private Plugin Hook Function
********************************

The private plugin hook function should bear the same name as the public plugin
hook function, but with a lower case first letter. The signature should match
the public plugin hook function, except that the first argument referencing the
PluginLoader should be removed. It should invoke the public plugin hook
function on ``plugins.DefaultPluginLoader``. It should always verify that the
DefaultPluginLoader is non-nil, log warning and return if the
DefaultPluginLoader has not been initialized.

In-Line Invocation
******************

Within the Geth codebase, the private plugin hook function should be invoked
with the appropriate arguments in a single line, to minimize unexpected
conflicts merging the upstream geth codebase into plugeth.
