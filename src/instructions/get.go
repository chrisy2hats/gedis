package instructions

import (
	"errors"
	"fmt"
	"kv"
	"strings"
)

type Get struct {
	Key   string
	value string
}

// Get instructions should be like "get foobar"
func (get *Get) Parse(raw_instruction string) error {
	fmt.Println("Parse get")
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return errors.New("Invalid get command")
	}
	splt := strings.Split(raw_instruction, " ")
	get.Key = splt[1]
	return nil
}

func (get *Get) Execute() (kv.MapValue, error) {

	if value, ok := kv.KV[get.Key]; ok {
		return value, nil
	}

	return kv.MapValue{}, errors.New("No Key " + get.Key)

}
