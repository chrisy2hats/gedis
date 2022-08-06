package instructions

import (
	"errors"
	"fmt"
	"kv"
	"regexp"
	"strings"
)

type Instruction interface {
	Execute() (kv.MapValue, error)
	Parse(raw_instruction string) error
}

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
