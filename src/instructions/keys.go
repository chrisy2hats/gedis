package instructions

import (
	"errors"
	"kv"
	"regexp"
	"strings"
)

type Keys struct {
	Key string
}

func ParseKeys(rawInstruction string) (Instruction, error) {
	spaceCount := strings.Count(rawInstruction, " ")
	if spaceCount != 1 {
		return nil, errors.New("Invalid keys command: " + rawInstruction)
	}
	splt := strings.Split(rawInstruction, " ")
	return &Keys{Key: splt[1]}, nil
}

func (Keys *Keys) Execute() (kv.MapValue, error) {
	// The startline and endline additons are needed to ensure the regex "a" doesn't match
	// both Keys called "a" and "about"
	regex := regexp.MustCompile("^" + Keys.Key + "$")
	bob := strings.Builder{}

	for k := range kv.KV {
		if regex.Match([]byte(k)) {
			bob.WriteString(k + "\n")
		}
	}
	return kv.MapValue{Value: bob.String()}, nil
}
