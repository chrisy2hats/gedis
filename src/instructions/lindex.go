package instructions

import (
	"errors"
	"fmt"
	"kv"
	"strconv"
	"strings"
)

type Lindex struct {
	Key   string
	Index int
}

func ParseLindex(rawInstruction string) (Instruction, error) {
	spaceCount := strings.Count(rawInstruction, " ")
	if spaceCount != 2 {
		return nil, errors.New("Invalid lindex command: " + rawInstruction)
	}
	splt := strings.Split(rawInstruction, " ")

	idx, err := strconv.Atoi(splt[2])
	if err != nil {
		return nil, err
	}

	return &Lindex{Key: splt[1], Index: idx}, nil
}

func (lindex *Lindex) Execute() (kv.MapValue, error) {
	if existingValue, exists := kv.KV[lindex.Key]; exists {
		if existingList, ok := existingValue.Value.([]string); ok {
			if lindex.Index < 0 || lindex.Index > len(existingList) {
				errMsg := fmt.Sprintf("%s %d %s %d",
					"Attempted to lindex outside of list range. Index provided:", lindex.Index, "Length of list:", len(existingList))
				return kv.MapValue{Value: ""}, errors.New(errMsg)
			}
			return kv.MapValue{Value: existingList[lindex.Index]}, nil
		} else {
			return kv.MapValue{Value: ""}, errors.New("lindex on a non list type key")
		}
	}

	return kv.MapValue{Value: ""}, nil
}
