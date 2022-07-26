package main

import (
	. "./instructions"
    "./config"
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

)

func parseInstruction(raw_instruction string) (interface{}, error) {
	without_newline := strings.TrimSpace(raw_instruction)

	if strings.HasPrefix(without_newline, "get") {
		instruction := Get{}
		err := instruction.Parse(without_newline)
		return instruction, err
	} else if strings.HasPrefix(without_newline, "set") {
		instruction := Set{}
		err := instruction.Parse(without_newline)
		return instruction, err
	} else if strings.HasPrefix(without_newline, "keys") {
		instruction := Keys{}
		err := instruction.Parse(without_newline)
		return instruction, err
	}
	return nil, errors.New("Unknown instruction: " + without_newline)
}

func HandleConnection(c net.Conn) {
    fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
	for {
        buffer, err := bufio.NewReader(c).ReadBytes('\n')
        if err != nil {
            fmt.Println("err", err)
            break
        }

        instruction, err := parseInstruction(string(buffer))
        if err != nil {
            fmt.Println("Failed to parse instruction: " + err.Error())
        }

        switch instruct := instruction.(type) {
        case Set:
            fmt.Println("set KEY:", instruct.Key)
            result, err := instruct.Execute()
            if err != nil {
                fmt.Println("err in get" + err.Error())
            }
            fmt.Println("Get result:", result)
        case Get:
            fmt.Println("get KEY:", instruct.Key)
            result, err := instruct.Execute()
            if err != nil {
                fmt.Println("err in get" + err.Error())
            }
            fmt.Println("Get result:", result)
            switch data := result.Value.(type) {
                case string:
                c.Write([]byte(data))        
                default:
                    errors.New("Unexpected type in Value field")
            }
        case Keys:
            fmt.Println("get KEY:", instruct.Key)
            result, err := instruct.Execute()
            if err != nil {
                fmt.Println("err in Keys" + err.Error())
            }
            fmt.Println("Get result:", result)
            c.Write([]byte(result))

        default:
            errors.New("Unknown instruction type")
        }
	}
}

func main() {
	fmt.Println("Gedis started")
    conn_str := config.CONN_HOST+":"+config.CONN_PORT
	l, err := net.Listen(config.CONN_TYPE, conn_str)
	if err != nil {
		fmt.Println("Failed to listen on : ",conn_str)
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
		go HandleConnection(c)
	}

}
