package main

import (
	"bufio"
	"config"
	"errors"
	"fmt"
	"instructions"
	"net"
	"os"
	"strings"
)

func parseInstruction(raw_instruction string) (instructions.Instruction, error) {
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
		instruction := instructions.Get{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "set") {
		instruction := instructions.Set{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "keys") {
		instruction := instructions.Keys{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	} else if strings.EqualFold(command, "lpush") {
		instruction := instructions.Lpush{}
		err := instruction.Parse(without_newline)
		return &instruction, err
	}

	// TODO proper return object here
	return &instructions.Get{}, errors.New("Unknown instruction: " + without_newline)
}

func handleConnection(c net.Conn) {
	for {
		buffer, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			fmt.Println("err", err)
			break
		}

		instruction, err := parseInstruction(string(buffer))
		if err != nil {
			fmt.Println("Failed to parse instruction: " + err.Error())
			continue
		}

		result, err := instruction.Execute()
		if err != nil {
			fmt.Println("Err in execution")
			continue
		}

		if result.Value != "" {
			switch res := result.Value.(type) {
			case string:
				fmt.Println("Get result:", res)
				c.Write([]byte(res))
			case []string:
				fmt.Println("Got a list result", res)
				//TODO
				//c.Write([]byte(res))
			default:
				errors.New("Unexpected map Value type")
			}
		}
	}
}

func main() {
	fmt.Println("Gedis started")
	conn_str := config.CONN_HOST + ":" + config.CONN_PORT
	l, err := net.Listen(config.CONN_TYPE, conn_str)
	if err != nil {
		fmt.Println("Failed to listen on : ", conn_str)
		os.Exit(1)
	}

	fmt.Println("Listening on " + conn_str)
	defer l.Close()
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			continue
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c)
	}

}
