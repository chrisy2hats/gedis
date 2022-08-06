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

func handleConnection(c net.Conn) {
	for {
		buffer, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			fmt.Println("err", err)
			break
		}

		instruction, err := instructions.ParseInstruction(string(buffer))
		if err != nil {
			fmt.Println("Failed to parse instruction: " + err.Error())
			continue
		}

		result, err := instruction.Execute()
		if err != nil {
			fmt.Println("Err in execution" + err.Error())
			// TODO report error to use
			continue
		}

		if result.Value != "" {
			switch res := result.Value.(type) {
			case string:
				fmt.Println("Get result:", res)
				c.Write([]byte(res))
			case []string:
				fmt.Println("Got a list result", res)
				c.Write([]byte(strings.Join(res, " ")))
			default:
				errors.New("Unexpected map Value type")
			}
		}
	}
}
