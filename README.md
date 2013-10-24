goutil
======

To start we now have 2 versions of a Map function. 
This allows us to iterate over Slice, Array and Map types
and execute a function on each of the elements in those
structures.

We can also run a CMap version of this function
which will loop and run the function concurrently
using go routines.
