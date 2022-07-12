package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

var m map[string]string

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

type Instruction struct {
	key string
}

type Get struct {
	Instruction
	value string
}

// Get instructions should be like "get foobar"
func (get *Get) parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return errors.New("Invalid get command")
	}
	splt := strings.Split(raw_instruction, " ")
	get.key = splt[1]
	return nil
}

func (get *Get) execute() (string, error) {
	if value, ok := m[get.key]; ok {
		return value, nil
	} else {
		return "", errors.New("No key " + get.key)
	}
}

type Set struct {
	Instruction
	value string
}

// Set instructions should be like "set foobar 42"
func (set *Set) parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 2 {
		return errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	set.key = splt[1]
	set.value = splt[2]
	return nil
}

func (set *Set) execute() (string, error) {
	m[set.key] = set.value
	return "", nil
}

type Keys struct {
	Instruction
	value string
}

func (set *Keys) parse(raw_instruction string) error {
	space_count := strings.Count(raw_instruction, " ")
	if space_count != 1 {
		return errors.New("Invalid set command")
	}
	splt := strings.Split(raw_instruction, " ")
	set.key = splt[0]
	set.value = splt[1]
	return nil
}

func (keys *Keys) execute() (string, error) {
	// The startline and endline additons are needed to ensure the regex "a" doesn't match
	// both keys called "a" and "about"
	regex := regexp.MustCompile("^" + keys.value + "$")
	str_builder := strings.Builder{}

	for k, _ := range m {
		if regex.Match([]byte(jk)) {
			str_builder.WriteString(k + "\n")
		}
	}
	return str_builder.String(), nil
}

func ParseInstruction(raw_instruction string) (interface{}, error) {
	without_newline := strings.TrimSpace(raw_instruction)

	if strings.HasPrefix(without_newline, "get") {
		instruction := Get{}
		err := instruction.parse(without_newline)
		return instruction, err
	} else if strings.HasPrefix(without_newline, "set") {
		instruction := Set{}
		err := instruction.parse(without_newline)
		return instruction, err
	} else if strings.HasPrefix(without_newline, "keys") {
		instruction := Keys{}
		err := instruction.parse(without_newline)
		return instruction, err
	}
	return nil, errors.New("Unknown instruction: " + without_newline)
}

func HandleConnection(c net.Conn) {
	for {
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		for {
			buffer, err := bufio.NewReader(c).ReadBytes('\n')
			if err != nil {
				fmt.Println("err", err)
				break
			}
			fmt.Println(string(buffer))
			instruction, err := ParseInstruction(string(buffer))
			if err != nil {
				fmt.Println("Failed to parse instruction: " + err.Error())
			}

			switch instruct := instruction.(type) {
			case Set:
				fmt.Println("set KEY:", instruct.key)
				result, err := instruct.execute()
				if err != nil {
					fmt.Println("err in get" + err.Error())
				}
				fmt.Println("Get result:", result)
				c.Write([]byte(result))
			case Get:
				fmt.Println("get KEY:", instruct.key)
				result, err := instruct.execute()
				if err != nil {
					fmt.Println("err in get" + err.Error())
				}
				fmt.Println("Get result:", result)
				c.Write([]byte(result))
			case Keys:
				fmt.Println("get KEY:", instruct.key)
				result, err := instruct.execute()
				if err != nil {
					fmt.Println("err in keys" + err.Error())
				}
				fmt.Println("Get result:", result)
				c.Write([]byte(result))

			default:
				errors.New("Unknown instruction type")

			}
		}
	}

}

func main() {
	fmt.Println("Gedis starting")
	m = make(map[string]string)

	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Failed to listen on port " + connPort)
		os.Exit(1)
	}
	defer l.Close()
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			continue
		}
		go HandleConnection(c)
	}

}
