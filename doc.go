/*
Package pooledioutil provides a re-implementation of some functions
from the ioutil package.

It uses an underlying `sync.Pool` for re-use of allocated buffers.
*/
package pooledioutil
