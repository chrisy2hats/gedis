package instructions

import (
	"kv"
	"testing"
)

func TestBadGetParsing(t *testing.T) {
	_, err := ParseGet("GET foobar 5 1")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseGet("GET dead beef 1")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestValidGetParsing(t *testing.T) {
	instruct, err := ParseGet("GET foobar")

	if get, ok := instruct.(*Get); ok {

		if get.Key != "foobar" {
			t.Error("Incorrect Get Key")
		}

	} else {
		t.Error("ParseGet returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{
		"foobar": kv.MapValue{Value: "5"},
	}

	instruct := Get{Key: "foobar"}

	result, err := instruct.Execute()
	if err != nil {
		t.Error(err.Error())
	}

	if (result != kv.MapValue{Value: "5"}) {
		t.Error("Unexpected result from Get exec")
	}

	if v, ok := kv.KV["foobar"]; ok {
		if val, ok := v.Value.(string); ok {
			if val != "5" {
				t.Error("Incorrect Value for key foobar")
			}
		} else {
			t.Error("Get key has incorrect type")
		}

	} else {
		t.Error("Key not in map after Set exec")
	}
}
