package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Instruction interface {
	Execute() (kv.MapValue, error)
	Parse(raw_instruction string) error
}

func ParseInstruction(raw_instruction string) (Instruction, error) {
	without_newline := strings.TrimSpace(raw_instruction)

	index := strings.Index(without_newline, " ")
	if index == -1 {
		index = len(without_newline)
	}

	command := without_newline[:index]

	// TODO ideally this string -> function can be stored in a map. Sadly the below doesn't work. Maybe there is a way to do similar
	//var _ = map[string]func(raw_instruction string) (instructions.Instruction, error){
	//	"get": instructions.ParseGet,
	//"set": instructions.Set.Parse,
	//}

	if strings.EqualFold(command, "get") {
		instruction := Get{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "set") {
		instruction := Set{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "keys") {
		instruction := Keys{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "lpush") {
		instruction := Lpush{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "lrange") {
		instruction := Lrange{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "lindex") {
		instruction := Lindex{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	}

	// TODO proper return object here
	return &Get{}, errors.New("Unknown instruction: " + without_newline)
}
