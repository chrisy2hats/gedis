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

func (get *Get) Execute() (kv.MapValue, error) {

	if value, ok := kv.KV[get.Key]; ok {
		return value, nil
	}

	return kv.MapValue{}, errors.New("No Key " + get.Key)
}

func ParseGet(raw_instruction string) (Instruction, error) {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return nil, errors.New("Invalid get command")
	}
	splt := strings.Split(raw_instruction, " ")
	return &Get{Key: splt[1]}, nil
}
