.. _install:

=======
Install 
=======

.. note:: Prior to install make sure to be familiar with :ref:`system requirements<system_req>`.


PluGeth can be installed in two ways. The repositories can be cloned and compiled from the source code. Alternatively PluGeth provides binaries of a PluGeth node as well as plugins.

In order to run PluGeth without the source code,  download the latest `release`_ here. 

The curated list of plugin builds can be found `here`_

.. note:: Make sure versions of PluGeth and plugins are compatable see: :ref:`version control<version>`.

After downloading plugins, move the ``.so`` files into the ``~/.ethereum/plugins`` directory. 

.. note:: The above location may change when changing ``--datadir``.








.. _release: https://github.com/openrelayxyz/plugeth/releases
.. _here: https://github.com/openrelayxyz/plugeth-plugins/releases