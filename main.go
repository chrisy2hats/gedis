package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

var m map[string]string

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

type Instruction struct {
	instruction string
	key         string
	arg         string
}

func ParseInstruction(raw_instruction string) (Instruction, error) {

	splt := strings.Fields(raw_instruction)
	if len(splt) < 3 {
		return Instruction{}, errors.New("Unknown command" + splt[0])
	}
	in := Instruction{splt[0], splt[1], splt[2]}
	fmt.Println(in)
	return in, nil
}

func HandleInstruction(instruction Instruction) (string, error) {
	if instruction.instruction == "set" {
		m[instruction.key] = instruction.arg
		return "", nil
	} else if instruction.instruction == "get" {
		return m[instruction.key], nil
	}

	return "", errors.New("Unknown instruction " + instruction.instruction)
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
			instr, err := ParseInstruction(string(buffer))
			if err != nil {
				fmt.Println(err)
			}
			val, err := HandleInstruction(instr)
			fmt.Println(m)
			fmt.Println("val", val)
			if val != "" {
				c.Write([]byte(val))
			}
		}
	}

}

func main() {
	m = make(map[string]string)
	fmt.Println("asd")
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
