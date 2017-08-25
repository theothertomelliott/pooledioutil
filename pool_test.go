package pooledioutil

import (
	"io/ioutil"
	"strings"
	"testing"
)

var testData = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func TestReadAll(t *testing.T) {
	pool := NewPool()
	reader := strings.NewReader(testData)
	read, err := pool.ReadAll(reader)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if string(read) != testData {
		t.Error("Output was not as expected")
	}
}

func TestReadFile(t *testing.T) {
	pool := NewPool()
	got, err := pool.ReadFile("testdata/lorem.txt")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	expected, err := ioutil.ReadFile("testdata/lorem.txt")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if string(expected) != string(got) {
		t.Error("Output was not as expected")
	}
}

func TestReadFileErr(t *testing.T) {
	pool := NewPool()
	_, err := pool.ReadFile("testdata/missing.txt")
	if err == nil {
		t.Error("Expected an error: ", err)
	}
}

func BenchmarkPooledReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool()
		for i := 0; i < 100; i++ {
			reader := strings.NewReader(testData)
			read, err := pool.ReadAll(reader)
			if err != nil {
				b.Error("Unexpected error: ", err)
			}
			if string(read) != testData {
				b.Error("Output was not as expected")
			}
		}
	}
}

func BenchmarkIoutilReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			reader := strings.NewReader(testData)
			read, err := ioutil.ReadAll(reader)
			if err != nil {
				b.Error("Unexpected error: ", err)
			}
			if string(read) != testData {
				b.Error("Output was not as expected")
			}
		}
	}
}

func BenchmarkPooledReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool := NewPool()
		for i := 0; i < 100; i++ {
			read, err := pool.ReadFile("testdata/lorem.txt")
			if err != nil {
				b.Error("Unexpected error: ", err)
			}
			if string(read) != testData {
				b.Error("Output was not as expected")
			}
		}
	}
}

func BenchmarkIoutilReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			read, err := ioutil.ReadFile("testdata/lorem.txt")
			if err != nil {
				b.Error("Unexpected error: ", err)
			}
			if string(read) != testData {
				b.Error("Output was not as expected")
			}
		}
	}
}
