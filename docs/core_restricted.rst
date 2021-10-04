.. _core_restricted:

=========================
PluGeth-utils Subpackages
=========================

PluGeth-utils is separated into two main packages: core, and restricted.

The `core` package has been implemented by the Rivet team, and is licensed under the MIT license, allowing it to be used in open source and closed source plugins alike. Nearly all plugins will need to import plugeth-utils/core in order to

The `restricted` package copies code from the go-ethereum project, which means it must be licensed under the LGPL license. If you import plugeth-utils/restricted, you must be sure that your plugin complies with requirements of linking to LGPL code, which will usually require making your source code available to anyone you distribute the plugin to.
