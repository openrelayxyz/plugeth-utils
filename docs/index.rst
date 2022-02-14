=======
PluGeth
=======



**The Geth fork to end all Forks.**

PluGeth is a fork of the Go Ethereum Client, `Geth`_, implementing the Golang plugin architecture allowing developers to adapt and extend Geth's capabilities using plugins rather than having to create additional new forks. 

The PluGeth project aims to provide a secure and versitile tool for anyone who needs to run a Geth (or Geth-derived) node client that supports features beyond those offered by Geth’s vanilla EVM. 

All dependencies and updates are handled by the PluGeth project, and so, PluGeth enables developers to focus on their projects without having to maintian upstream code.  


- :ref:`project`
- :ref:`install`
- :ref:`build`     
- :ref:`custom`

.. toctree::
    :maxdepth: 1
    :caption: Overview
    :hidden:

    project
    types
    

.. toctree::
    :maxdepth: 1
    :caption: Tutorials
    :hidden:

    install
    build
    custom
    

.. toctree::
    :maxdepth: 1
    :caption: Reference
    :hidden:

    existing
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
    :hidden:

    contact








.. _Geth: https://geth.ethereum.org/