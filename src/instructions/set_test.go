package instructions

import (
	"kv"
	"testing"
)

func TestBadSetParsing(t *testing.T) {
	_, err := ParseSet("SET foobar 5 1")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseSet("SET foo bar 1")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestValidSetParsing(t *testing.T) {
	instruct, err := ParseSet("SET foobar 5")

	if set, ok := instruct.(*Set); ok {
		if set.Value != "5" {
			t.Error("Incorrect Set Value")
		}

		if set.Key != "foobar" {
			t.Error("Incorrect Set Key")
		}

	} else {
		t.Error("ParseSet returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestSetExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{}

	instruct := Set{Key: "foobar", Value: "1"}

	result, err := instruct.Execute()
	if err != nil {
		t.Error(err.Error())
	}

	if (result != kv.MapValue{}) {
		t.Error("Unexpected result from Set exec")
	}

	if v, ok := kv.KV["foobar"]; ok {
		if val, ok := v.Value.(string); ok {
			if val != "1" {
				t.Error("Incorrect Value for key foobar")
			}
		} else {
			t.Error("Set key has incorrect type")
		}

	} else {
		t.Error("Key not in map after Set exec")
	}
}
