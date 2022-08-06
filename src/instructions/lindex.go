package instructions

import (
	"errors"
	"kv"
	"strconv"
	"strings"
)

type Lindex struct {
	Key   string
	index int
}

func (lindex *Lindex) Parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return errors.New("Invalid lindex command: " + raw_instruction)
	}
	splt := strings.Split(raw_instruction, " ")

	lindex.Key = splt[1]
	index, err := strconv.Atoi(splt[2])
	if err != nil {
		return err
	}
	lindex.index = index

	return nil
}

func (lindex *Lindex) Execute() (kv.MapValue, error) {
	if existing_value, exists := kv.KV[lindex.Key]; exists {
		if existing_list, ok := existing_value.Value.([]string); ok {
			if lindex.index < 0 || lindex.index > len(existing_list) {
				return kv.MapValue{Value: ""}, errors.New("User tried to lindex outside of list range")
			}

			return kv.MapValue{Value: existing_list[lindex.index]}, nil
		} else {
			return kv.MapValue{Value: ""}, errors.New("User tried to lpush a non list type")
		}
	}

	return kv.MapValue{Value: ""}, nil
}
