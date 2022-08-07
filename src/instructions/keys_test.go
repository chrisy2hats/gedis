package instructions

import (
	"kv"
	"reflect"
	"testing"
)

func TestBadKeysParsing(t *testing.T) {
	_, err := ParseKeys("KEYS a 5")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseKeys("KEYS")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestValidKeysParsing(t *testing.T) {
	instruct, err := ParseKeys("KEYS *")

	if keys, ok := instruct.(*Keys); ok {

		if keys.Key != "*" {
			t.Error("Incorrect Get")
		}

	} else {
		t.Error("ParseKeys returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestKeysExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{
		"foobar": kv.MapValue{Value: "1"},
		"foo":    kv.MapValue{Value: "2"},
		"bar":    kv.MapValue{Value: "3"},
		"a":      kv.MapValue{Value: "3"},
	}

	expected := map[string]string{
		"a":     "a\n",
		"*bar*": "foobar\nbar\n",
	}
	for k, v := range expected {
		instruct := Keys{Key: k}
		result, err := instruct.Execute()
		if err != nil {
			t.Error(err.Error())
		}

		if val, ok := result.Value.(string); ok {
			if val != v {
				t.Error("Unexpected result from Keys exec whilst processing: " + v + ":" + val)
			}
		} else {
			t.Error("Incorrect type returned from KEYS" + reflect.TypeOf(result.Value).String())
		}
	}

}
