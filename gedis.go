package main

import (
	"bufio"
	"config"
	"encoding/json"
	"fmt"
	"instructions"
	"net"
	"os"
	"sync"
)

func main() {
	fmt.Println("Gedis started")
	connStr := config.CONN_HOST + ":" + config.CONN_PORT
	l, err := net.Listen(config.CONN_TYPE, connStr)
	if err != nil {
		fmt.Println("Failed to listen on : ", connStr)
		os.Exit(1)
	}

	fmt.Println("Listening on " + connStr)
	defer l.Close()

	lock := sync.Mutex{}
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			continue
		}
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c, &lock)
	}
}

func handleConnection(c net.Conn, lock *sync.Mutex) {
	for {
		var outcome = map[string]interface{}{
			"successful": true,
		}

		buffer, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			outcome["successful"] = false
			outcome["error"] = err.Error()

			if err.Error() == "EOF" {
				break
			}
		}

		instruction, err := instructions.ParseInstruction(string(buffer))
		if err != nil {
			outcome["successful"] = false
			outcome["error"] = err.Error()
			jsn, _ := json.Marshal(outcome)
			_, err := c.Write(jsn)
			if err != nil {
				fmt.Println("Write to socket failed. Client: " + c.RemoteAddr().String() + " Error: " + err.Error())
			}
			continue
		}

		// All commands require the entire map to be locked
		// Commands that may add a key like SET, LPUSH can't be done concurrently
		// And commands that only read a key can't either as a different thread could delete a key whilst a GET is processed
		lock.Lock()
		result, err := instruction.Execute()
		lock.Unlock()

		if err != nil {
			outcome["successful"] = false
			outcome["error"] = err.Error()
			jsn, _ := json.Marshal(outcome)
			_, err := c.Write(jsn)
			if err != nil {
				fmt.Println("Write to socket failed. Client: " + c.RemoteAddr().String() + " Error: " + err.Error())
			}
			continue
		}

		if result.Value != "" && result.Value != nil {
			outcome["result"] = result.Value
		}
		jsn, _ := json.Marshal(outcome)
		_, err = c.Write(jsn)
		if err != nil {
			fmt.Println("Write to socket failed. Client: " + c.RemoteAddr().String() + " Error: " + err.Error())
			continue
		}
	}
}
