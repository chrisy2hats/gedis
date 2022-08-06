package instructions

import (
	"kv"
)

type Instruction interface {
	Execute() (kv.MapValue, error)
	Parse(raw_instruction string) error
}
