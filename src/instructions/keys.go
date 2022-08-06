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

func (set *Keys) Parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	set.Key = splt[0]
	set.value = splt[1]
	return nil
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
