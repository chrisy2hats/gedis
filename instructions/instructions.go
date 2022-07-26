package instructions

import (
	. "../map"
	"errors"
	"regexp"
	"strings"
)

type Instruction struct {
	Key string
}

type Get struct {
	Instruction
	value string
}

type Set struct {
	Instruction
	value string
}

type Keys struct {
	Instruction
	value string
}

// Get instructions should be like "get foobar"
func (get *Get) Parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return errors.New("Invalid get command")
	}
	splt := strings.Split(raw_instruction, " ")
	get.Key = splt[1]
	return nil
}

func (get *Get) Execute() (MapValue, error) {
    
	if value, ok := KV[get.Key]; ok {
		return value, nil
    }

    return MapValue{}, errors.New("No Key " + get.Key)
	
}

// Parse Set instructions which are like "set foobar 42"
func (set *Set) Parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	set.Key = splt[1]
	set.value = splt[2]
	return nil
}

func (set *Set) Execute() (MapValue, error) {
	KV[set.Key] = MapValue{Value:set.value}
	return MapValue{}, nil
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

func (Keys *Keys) Execute() (string, error) {
	// The startline and endline additons are needed to ensure the regex "a" doesn't match
	// both Keys called "a" and "about"
	regex := regexp.MustCompile("^" + Keys.value + "$")
	str_builder := strings.Builder{}

	for k := range KV {
		if regex.Match([]byte(k)) {
			str_builder.WriteString(k + "\n")
		}
	}
	return str_builder.String(), nil
}
