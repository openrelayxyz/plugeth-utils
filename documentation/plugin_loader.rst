.. _plugin_loader:

=============
Plugin Loader
=============

.. todo:: Ausitn, take a pass at flushing this out. 

At the heart of the PluGeth project is the `PluginLoader`_. 

Upon invocation the PluginLoader will parse through a list of known plugins and either return the plugin name passed to it or, if not found, append to that list. Additionally the loader will check the function signature of the plugin to assure complience with anticipated behavior. Once these checks are passed and the plugin name and function signature is validated the plugin will be invoked.  



.. _PluginLoader: https://github.com/openrelayxyz/plugeth/blob/develop/plugins/plugin_loader.go