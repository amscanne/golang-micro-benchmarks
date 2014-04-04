Golang Benchmarks
=================

This is a tiny collection of micro benchmarks.

The intent is to compare some language feature to others.

The original benchmark arose from a simple question:

Suppose a runtime parameter is dynamic, but generically fixed.

* Is it best to build a new closure when it changes?
* Or simply evaluate the parameter in an if, continuously?
* Or should we write generic types, and use an interface?

Running
-------

To run the suite, use:

    go test -bench=.
