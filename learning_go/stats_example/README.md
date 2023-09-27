## Stats Example

This is a simple Go program that prints the linear regressions series of the
data sets in [Anscombe's Quartet](https://en.wikipedia.org/wiki/Anscombe%27s_quartet)


### Installation
---

#### Go install

The program can be installed by running `go install stats_example.go` from
within this directory.

#### Makefile

A makefile is included with the project and will create executable binaries for
Linux, OSX, and Windows that will be stored to the bin directory. An example of
the workflow on OSX would be:
```
cd path/to/this/directory
make
./bin/anscombe-quartet-darwin
...
```
