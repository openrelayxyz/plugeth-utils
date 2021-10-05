.. _version:

=======
Version
=======

PluGeth is separated into three packages in order to minimize dependency conflicts. Golang plugins cannot include different versions of the same packages as the program loading the plugin. If plugins had to import packages from PluGeth itself, a plugin build could only be loaded by that same version of PluGeth. By separating out the PluGeth-utils package, both PluGeth and the plugins must rely on the same version of PluGeth-utils, but plugins can be compatible with any version of PluGeth compiled with the same version of PluGeth-utils.

PluGeth builds will follow the naming convention:

.. code-block:: shell

   geth-$PLUGETH_UTILS_VERSION-$GETH_VERSION-$RELEASE

For example:

.. code-block:: shell

   geth-0.1.0-1.10.8-0

Tells us that:

* PluGeth-utils version is 0.1.0
* Geth version is 1.10.8
* This is the first release with that combination of dependencies.

Plugin builds will follow the naming convention:

.. code-block:: shell

   $PLUGIN_NAME-$PLUGETH_UTILS_VERSION-$PLUGIN_VERSION

For example:

.. code-block:: shell

   blockupdates-0.1.0-1.0.2

Tells us that:

* The plugin is "blockupdates"
* The PluGeth-utils version is 0.1.0
* The plugin version is 1.0.2

When a Geth update comes out, you can expect a release of `geth-0.1.0-1.10.9-0`, which will be compatible with the same set of plugins.

When PluGeth upgrades are necessary, plugins will need to be recompiled. Whenever possible, we will try to avoid forcing plugins to be recompiled for an immediate Geth upgrade. For example, if we have geth-0.1.0-1.10.8, and upgrade PluGeth-utils, we will have a geth-0.1.1-1.10.8, followed by a geth-0.1.1-1.10.9. This will give users time to upgrade plugins from PluGeth-utils 0.1.0 to 0.1.1 while staying on Geth 1.10.8, and when it is time to upgrade to Geth 1.10.9 they can continue using the plugins they were using with geth 1.10.8. Depending on upgrades to Geth, it may not always be possible to maintain compatibility with existing PluGeth versions, which will be noted in release notes.
