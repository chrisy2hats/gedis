package kv

import (
	"sync"
)

type MapValue struct {
	// Can be a string for simple key value pairs or an array for key list pairs
	Value interface{}

	// Go slices are not thread safe so appending to a key list pair requires a lock
	// https://pkg.go.dev/reflect#SliceHeader
	writeLock sync.RWMutex
}

var KV = make(map[string]MapValue)
