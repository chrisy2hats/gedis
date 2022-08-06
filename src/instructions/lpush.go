package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Lpush struct {
	Key   string
	value string
}

func (lpush *Lpush) Parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return errors.New("Invalid lpush command")
	}
	splt := strings.Split(raw_instruction, " ")
	lpush.Key = splt[1]
	lpush.value = splt[2]
	return nil
}

func (lpush *Lpush) Execute() (kv.MapValue, error) {
	if existing_value, exists := kv.KV[lpush.Key]; exists {
		if existing_list, ok := existing_value.Value.([]string); ok {
			// TODO there must be a better way to do this...
			kv.KV[lpush.Key] = kv.MapValue{Value: append(existing_list, lpush.value)}
		} else {
			return kv.MapValue{Value: ""}, errors.New("User tried to lpush a non list type")
		}

	} else {
		kv.KV[lpush.Key] = kv.MapValue{Value: []string{lpush.value}}
	}

	return kv.MapValue{Value: ""}, nil
}