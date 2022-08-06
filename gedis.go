package main

import (
	"bufio"
	"config"
	"encoding/json"
	"fmt"
	"instructions"
	"net"
	"os"
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

		var outcome = map[string]interface{}{
			"successful": true,
		}

		instruction, err := instructions.ParseInstruction(string(buffer))
		if err != nil {
			outcome["successful"] = false
			outcome["error"] = err.Error()
			jsn, _ := json.Marshal(outcome)
			c.Write(jsn)
			continue
		}

		result, err := instruction.Execute()
		if err != nil {
			fmt.Println("Err in execution" + err.Error())
			outcome["successful"] = false
			outcome["error"] = err.Error()
			jsn, _ := json.Marshal(outcome)
			c.Write(jsn)
			continue
		}

		if result.Value != "" && result.Value != nil {
			outcome["result"] = result.Value
		}
		jsn, _ := json.Marshal(outcome)
		c.Write(jsn)
	}
}
