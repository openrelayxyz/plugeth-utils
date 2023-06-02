.. _plugin_loader:

=============
Plugin Loader
=============

The Plugin Loader is provided to each Plugin through the Initialize()``
function. It provides plugins with:


Lookup
======
``Lookup(name string, validate func(interface{}) bool) []interface{}``

Returns a list of values from plugins identified by ``name``, which match the
provided ``validate`` predicate. For example:


.. code-block:: go

    pl.Lookup("Version", func(item interface{}) bool {
      _, ok := item.(int)
      return ok
    })

Would return a list of ``int`` objects named ``Version`` in any loaded plugins.
This can enable Plugins to interact with each other, accessing values and
functions implemented in other plugins.

GetFeed
=======
``GetFeed() Feed``

Returns a new feed that the plugin can use for publish/subscribe models.

For example:

.. code-block:: go

    feed := pl.GetFeed()
    go func() {
      ch := make(chan string)
      sub := feed.Subscribe(ch)
      for {
        select {
        case item := <-ch:
          // Do something with item
        case err := <sub.Err():
          log.Error("An error has occurred", "err", err)
          sub.Unsubscribe()
          close(ch)
          return
        }
      }
    }()

    feed.Send("hello")
    feed.Send("world")


Note that you can send any type through a feed, but the subscribed channel and
sent objects must be of matching types.
