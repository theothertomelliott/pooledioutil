package pooledioutil

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// Pool provides a pool for re-using memory allocated for Read operations.
type Pool struct {
	bufferPool sync.Pool
}

// NewPool creates an initialized pool
func NewPool() *Pool {
	return &Pool{
		bufferPool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// ReadAll reads from r until an error or EOF and returns the data it read.
func (p *Pool) ReadAll(r io.Reader) (b []byte, err error) {
	buf := p.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer p.bufferPool.Put(buf)
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

// ReadFile reads the file named by filename and returns the contents.
//
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func (p *Pool) ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return p.ReadAll(f)
}

var sharedPool = NewPool()

// ReadAll reads from r until an error or EOF and returns the data it read.
//
// This is a wrapper around *Pool.ReadAll, using a single global pool.
func ReadAll(r io.Reader) (b []byte, err error) {
	return sharedPool.ReadAll(r)
}

// ReadFile reads the file named by filename and returns the contents.
//
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
//
// This is a wrapper around *Pool.ReadFile, using a single global pool.
func ReadFile(filename string) ([]byte, error) {
	return sharedPool.ReadFile(filename)
}
