# pooledioutil

[![GoDoc](https://godoc.org/github.com/theothertomelliott/pooledioutil?status.svg)](https://godoc.org/github.com/theothertomelliott/pooledioutil)

Package pooledioutil provides a re-implementation of some functions
from the ioutil package.

It uses an underlying `sync.Pool` for re-use of allocated buffers.

Benchmark results against ioutil functions performing 100 sequential reads.

    BenchmarkPooledReadAll-4          100000             18367 ns/op            6125 B/op        105 allocs/op
    BenchmarkIoutilReadAll-4           30000             46411 ns/op          208000 B/op        300 allocs/op
    BenchmarkPooledReadFile-4           3000            387648 ns/op           10125 B/op        305 allocs/op
    BenchmarkIoutilReadFile-4           3000            462530 ns/op          130400 B/op        500 allocs/op