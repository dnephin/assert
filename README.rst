
assert
======

``assert`` is an golang assertions library for unit testing.


Why assert?
-----------

The ``testing`` package lease a lot to be desired from a testing framework.
Using ``t.Fatal()`` for assertions leads to a few problems:

* tests are less readable when they are covered with ``if ... { t.Fatal(...) }```
* failure messages are unlikely to be consistent across packages, and sometimes
  even within the same package
* failures messages may be wrong, misleading or woefully incomplete because it
  is up to the programmer to write a message that matches the assertion. If
  the assertion changes it is easy to forget to update the failure message.

``assert`` aims to solve these problems by providing a minimal library that
contains general purpose assertions that can be used across many packages,
and provide accurate, and helpful error messages.


Contributing
------------

To get started with ``assert``:

.. code::

    make shell
    glide install
    hack/watch


Related libraries
-----------------

``assert`` takes inspiration from `pytest <http://pytest.org/>`_ assertions,
where only the bare python ``assert`` statement is used, and the framework is
responsible for generating a helpful error message.

`stretchr/testify/assert <https://github.com/stretchr/testify#assert-package>`_
also provides a set of assertions, however they use ``t.Errorf()`` to report
an error instead of failing the test.
