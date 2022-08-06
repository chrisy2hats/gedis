package instructions

import (
	"errors"
	"kv"
	"regexp"
	"strings"
)

type Keys struct {
	Key   string
	value string
}

func ParseKeys(raw_instruction string) (Instruction, error) {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return nil, errors.New("Invalid keys command")
	}
	splt := strings.Split(raw_instruction, " ")
	return &Keys{Key: splt[0], value: splt[1]}, nil
}

func (Keys *Keys) Execute() (kv.MapValue, error) {
	// The startline and endline additons are needed to ensure the regex "a" doesn't match
	// both Keys called "a" and "about"
	regex := regexp.MustCompile("^" + Keys.value + "$")
	str_builder := strings.Builder{}

	for k := range kv.KV {
		if regex.Match([]byte(k)) {
			str_builder.WriteString(k + "\n")
		}
	}
	return kv.MapValue{Value: str_builder.String()}, nil
}
