package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Lpush struct {
	Key   string
	Value string
}

func ParseLpush(rawInstruction string) (Instruction, error) {
	spaceCount := strings.Count(rawInstruction, " ")
	if spaceCount != 2 {
		return nil, errors.New("Invalid lpush command")
	}
	splt := strings.Split(rawInstruction, " ")
	return &Lpush{Key: splt[1], Value: splt[2]}, nil
}

func (lpush *Lpush) Execute() (kv.MapValue, error) {
	if existingValue, exists := kv.KV[lpush.Key]; exists {
		if existingList, ok := existingValue.Value.([]string); ok {
			// TODO there must be a better way to do this...
			kv.KV[lpush.Key] = kv.MapValue{Value: append(existingList, lpush.Value)}
		} else {
			return kv.MapValue{Value: ""}, errors.New("lpush on a non list type key")
		}

	} else {
		kv.KV[lpush.Key] = kv.MapValue{Value: []string{lpush.Value}}
	}

	return kv.MapValue{}, nil
}
