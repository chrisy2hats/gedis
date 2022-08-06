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

func ParseInstruction(rawInstruction string) (Instruction, error) {
	withoutNewline := strings.TrimSpace(rawInstruction)

	index := strings.Index(withoutNewline, " ")
	if index == -1 {
		index = len(withoutNewline)
	}

	command := withoutNewline[:index]

	if f, exists := parsingLookup[strings.ToLower(command)]; exists {
		return f(withoutNewline)
	} else {
		return nil, errors.New("Unknown instruction: " + withoutNewline)
	}
}
