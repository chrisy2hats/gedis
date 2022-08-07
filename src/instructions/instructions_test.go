package instructions

import (
	"testing"
)

func TestNonExistentComm(t *testing.T) {
	intruction, err := ParseInstruction("FAKE foo 1")
	if err == nil {
		t.Error("No error on non existent COMMAND")
	}
	if intruction != nil {
		t.Error("Unexpected instruction. Expected nil")
	}

}

func TestSetParse(t *testing.T) {
	instruct, err := ParseInstruction("SET foobar 0")

	if set, ok := instruct.(*Set); ok {
		if set.Value != "0" {
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
