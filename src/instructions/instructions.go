package instructions

import (
	"errors"
	"kv"
	"strings"
)

type Instruction interface {
	Execute() (kv.MapValue, error)
}

var parsingLookup = map[string]func(string) (Instruction, error){
	"get":    ParseGet,
	"set":    ParseSet,
	"keys":   ParseKeys,
	"lpush":  ParseLpush,
	"lindex": ParseLindex,
	"lrange": ParseLrange,
}

func ParseInstruction(raw_instruction string) (Instruction, error) {
	without_newline := strings.TrimSpace(raw_instruction)

	index := strings.Index(without_newline, " ")
	if index == -1 {
		index = len(without_newline)
	}

	command := without_newline[:index]

	if f, exists := parsingLookup[strings.ToLower(command)]; exists {
		return f(without_newline)
	} else {
		return nil, errors.New("Unknown instruction: " + without_newline)
	}
}
