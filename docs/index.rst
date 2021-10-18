=======
PluGeth
=======

PluGeth is a fork of the Go Ethereum Client, `Geth`_, that implements a plugin architecture allowing developers to extend Geth's capabilities in a number of different ways using plugins rather than having to create additional new forks. 

From Here:
----------

- Ready for an  overview of the project and some context? :ref:`project`
- If your goal is to run existing plugns without sourcecode: :ref:`install`
- If your goal is to build and deploy existing plugins or make custom plugins: :ref:`build`     

- If your goal is to build cutsom plugins: :ref:`custom`

.. warning:: Right now PluGeth is in early development. We are 
             still settling on some of the plugin APIs, and are
             not yet making official releases. From an operational 
             perspective, PluGeth should be as stable as upstream Geth less whatever instability is added by plugins you might run. But if you plan to run PluGeth today, be aware that furture updates will likely break your plugins. 

Table of Contents
*****************


.. toctree::
    :maxdepth: 1
    :caption: Overview

    project
    types
    

.. toctree::
    :maxdepth: 1
    :caption: Tutorials

    install
    build
    custom
    

.. toctree::
    :maxdepth: 1
    :caption: Reference

    system_req
    version
    api
    plugin_loader
    hooks
    hook_writing
    core_restricted

.. toctree::
    :maxdepth: 1
    :caption: Contact

    contact



.. _Geth: https://geth.ethereum.org/