package instructions

import (
	"kv"
	"reflect"
	"testing"
)

func TestBadLpushParsing(t *testing.T) {
	_, err := ParseLpush("LPUSH foobar 1 5 ")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseLpush("LPUSH dead beef 1")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestValidLpushParsing(t *testing.T) {
	instruct, err := ParseLpush("LPUSH foobar 1")

	if lpush, ok := instruct.(*Lpush); ok {

		if lpush.Key != "foobar" {
			t.Error("Incorrect Lpush Key")
		}

		if lpush.Value != "1" {
			t.Error("Incorrect Lpush Key")
		}

	} else {
		t.Error("ParseLpush returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestLpushExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{
		"foobar": kv.MapValue{Value: []string{"foo", "bar"}},
	}

	instruct := Lpush{Key: "foobar", Value: "I want a snack"}

	result, err := instruct.Execute()
	if err != nil {
		t.Error(err.Error())
	}

	if (result != kv.MapValue{}) {
		t.Error("Unexpected result from Lpush exec")
	}

	if v, ok := kv.KV["foobar"]; ok {
		if val, ok := v.Value.([]string); ok {
			if !reflect.DeepEqual(val, []string{"foo", "bar", "I want a snack"}) {

				t.Error("Incorrect Value for key foobar")
			}
		} else {
			t.Error("Get key has incorrect type")
		}

	} else {
		t.Error("Key not in map after Set exec")
	}
}
