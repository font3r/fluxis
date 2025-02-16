package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const buffer_limit int = 1024

func main() {
	args := os.Args
	conn, err := net.Dial("tcp4", "0.0.0.0:5845")
	if err != nil {
		fmt.Println("dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(args[1]))
	if err != nil {
		fmt.Println("write to server failed ", err.Error())
		os.Exit(1)
	}

	buffer := make([]byte, buffer_limit)
	res := strings.Builder{}
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Printf("ERROR: error reading from connection - %s\n", err)
			}
			break
		}

		res.Write(buffer[:n])

		if n < buffer_limit {
			break
		}
	}

	fmt.Printf("> %s", res.String())
	conn.Close()
}
