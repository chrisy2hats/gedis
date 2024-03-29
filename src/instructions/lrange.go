package instructions

import (
	"errors"
	"kv"
	"math"
	"strconv"
	"strings"
)

type Lrange struct {
	Key   string
	Start int
	End   int
}

func ParseLrange(raw_instruction string) (Instruction, error) {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 3 {
		return nil, errors.New("Invalid lrange command: " + raw_instruction)
	}
	splt := strings.Split(raw_instruction, " ")

	start, err := strconv.Atoi(splt[2])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(splt[3])
	if err != nil {
		return nil, err
	}

	return &Lrange{Key: splt[1], Start: start, End: end}, nil
}

func (lrange *Lrange) Execute() (kv.MapValue, error) {
	if existing_value, exists := kv.KV[lrange.Key]; exists {
		if existing_list, ok := existing_value.Value.([]string); ok {
			start := int(math.Min(float64(lrange.Start), float64(len(existing_list))))
			end := int(math.Min(float64(lrange.End), float64(len(existing_list))))
			return kv.MapValue{Value: existing_list[start:end]}, nil
		} else {
			return kv.MapValue{Value: ""}, errors.New("User tried to lpush a non list type")
		}
	}

	return kv.MapValue{Value: ""}, nil
}
