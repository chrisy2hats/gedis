package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Set struct {
	Key   string
	value string
}

func ParseSet(rawInstruction string) (Instruction, error) {

	spaceCount := strings.Count(rawInstruction, " ")
	if spaceCount != 2 {
		return nil, errors.New("Invalid set command: " + rawInstruction)
	}
	splt := strings.Split(rawInstruction, " ")
	return &Set{Key: splt[1], value: splt[2]}, nil
}

func (set *Set) Execute() (kv.MapValue, error) {
	kv.KV[set.Key] = kv.MapValue{Value: set.value}
	return kv.MapValue{}, nil
}
