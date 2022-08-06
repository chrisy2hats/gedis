package kv

type MapValue struct {
	// Can be a string for simple key value pairs or an array for key list pairs
	Value interface{}
}

var KV = make(map[string]MapValue)
