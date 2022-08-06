package instructions

import (
	"errors"
	"fmt"
	"kv"
	"strings"
)

type Set struct {
	Key   string
	value string
}

func ParseSet(raw_instruction string) (Instruction, error) {

	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return nil, errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	return &Set{Key: splt[1], value: splt[2]}, nil
}

func (set *Set) Execute() (kv.MapValue, error) {
	fmt.Println("Setin", set.Key)
	kv.KV[set.Key] = kv.MapValue{Value: set.value}
	return kv.MapValue{}, nil
}
