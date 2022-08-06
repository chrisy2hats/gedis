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

// Parse Set instructions which are like "set foobar 42"
func (set *Set) Parse(raw_instruction string) error {
	fmt.Println("Parse set")

	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	set.Key = splt[1]
	set.value = splt[2]
	return nil
}

func (set *Set) Execute() (kv.MapValue, error) {
	fmt.Println("Setin", set.Key)
	kv.KV[set.Key] = kv.MapValue{Value: set.value}
	return kv.MapValue{}, nil
}
