.. Plugeth documentation master file, created by
   sphinx-quickstart on Tue Sep 21 16:08:24 2021.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.
=======
PluGeth
=======

PluGeth is a fork of the `Go Ethereum Client Geth`_ that implements a plugin architecture, allowing developers to extend Geth's capabilities in a number of different ways using plugins, rather than having to create additional, new forks of Geth.

.. warning:: Right now PluGeth is in early development. We are 
             still settling on some of the plugin APIs, and are
             not yet making official releases. From an operational 
             perspective, PluGeth should be as stable as upstream Geth less whatever isstability is added by plugins you might run. But if you plan to run PluGeth today, be aware that furture updates will likely break you plugins. 



Table of Contents
*****************


.. toctree::
    :maxdepth: 1
    :caption: Overview

    project
    types
    hooks
    anatomy
    

.. toctree::
    :maxdepth: 1
    :caption: Tutorials

    build

.. toctree::
    :maxdepth: 1
    :caption: Reference

    system_req
    plugin_loader
    plugin_hooks
    core_restricted
    api

.. toctree::
    :maxdepth: 1
    :caption: Contact

    contact


.. _Go Ethereum Client Geth: https://github.com/ethereum/go-ethereum