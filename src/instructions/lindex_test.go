package instructions

import (
	"kv"
	"testing"
)

func TestBadLindexParsing(t *testing.T) {
	_, err := ParseLindex("LINDEX foobar 0 5")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseLindex("LINDEX 5")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestValidLindexParsing(t *testing.T) {
	instruct, err := ParseLindex("LINDEX foobar 0")

	if lindex, ok := instruct.(*Lindex); ok {
		if lindex.Index != 0 {
			t.Error("Incorrect lindex Index")
		}
	} else {
		t.Error("ParseLindex returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestLindexExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{
		"foobar": kv.MapValue{Value: []string{"1", "a"}},
	}

	instruct := Lindex{Key: "foobar", Index: 0}

	result, err := instruct.Execute()
	if err != nil {
		t.Error(err.Error())
	}
	if result.Value != "1" {
		t.Error("Incorrect Value")
	}
}
