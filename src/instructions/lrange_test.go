package instructions

import (
	"kv"
	"reflect"
	"strings"
	"testing"
)

func TestBadLrangehParsing(t *testing.T) {
	_, err := ParseLrange("LRANGE foobar 10")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = ParseLrange("LRANGE foobar")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestValidLrangeParsing(t *testing.T) {
	instruct, err := ParseLrange("LRANGE foobar 0 5")

	if lrange, ok := instruct.(*Lrange); ok {

		if lrange.Key != "foobar" {
			t.Error("Incorrect Lrange Key")
		}

		if lrange.Start != 0 {
			t.Error("Incorrect Lrange start")
		}

		if lrange.End != 5 {
			t.Error("Incorrect Lrange end")
		}

	} else {
		t.Error("ParseLrange returned wrong type")
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func TestLrangeExec(t *testing.T) {
	kv.KV = map[string]kv.MapValue{
		"foobar": kv.MapValue{Value: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	expected := map[Lrange][]string{
		Lrange{Key: "foobar", Start: 0, End: 3}: {"a", "b", "c"},
	}

	for k, v := range expected {

		result, err := k.Execute()
		if err != nil {
			t.Error(err.Error())
		}

		if !reflect.DeepEqual(result.Value, []string{"a", "b", "c"}) {
			t.Error("Unexpected result from Lrange exec whilst processing: " + strings.Join(v, " "))
		}
	}

}
