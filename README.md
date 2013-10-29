goutil -> itertools
======

Map - function which allows us to iterate over Slice, Array, chan and Map types
      and execute a function on each of the elements in those
      structures.

CMap - version of the Map function
        which will loop and run the function concurrently
        using go routines.

Iterate - function which turns all typically
            iterable objects into object accepted by the range command

Range - function to create a generator
          similar to the xrange python function.

Pair - a type, which is
        in line with the c++ Pair object
