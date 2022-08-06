package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Get struct {
	Key   string
	value string
}

func ParseGet(rawInstruction string) (Instruction, error) {
	spaceCount := strings.Count(rawInstruction, " ")
	if spaceCount != 1 {
		return nil, errors.New("Invalid get command: " + rawInstruction)
	}
	splt := strings.Split(rawInstruction, " ")
	return &Get{Key: splt[1]}, nil
}

func (get *Get) Execute() (kv.MapValue, error) {

	if value, ok := kv.KV[get.Key]; ok {
		return value, nil
	}

	return kv.MapValue{}, errors.New("No Key " + get.Key)
}
