package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

const port = 5845
const buffer_limit int = 1024

func main() {
	fmt.Println("INFO: starting fluxis server")

	st := NewStorage()

	start(func(message string) string {
		if strings.HasPrefix(message, "SET") {
			cmd := strings.Split(message[3:], "=")
			err := st.SetKey(cmd[0], cmd[1])
			if err != nil {
				return "ERR"
			}

			return "OK"
		}

		if strings.HasPrefix(message, "GET") {
			return st.GetKey(message[3:])
		}

		if strings.HasPrefix(message, "DEBUG") {
			return st.Debug()
		}

		return "INVALID COMMAND"
	})
}

func start(process func(s string) string) {
	ln, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(fmt.Sprintf("ERROR: error during binding to port %d\n", port))
	}

	defer ln.Close()
	fmt.Printf("INFO: listening on %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("ERROR: during accepting connection from %v\n", err.Error())
			continue
		}

		fmt.Printf("DEBUG: accepted connection from %s\n", conn.RemoteAddr())

		req := strings.Builder{}
		buffer := make([]byte, buffer_limit)

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Printf("ERROR: error reading from connection - %s\n", err)
				}

				fmt.Println("DEBUG: end of input")
				break
			}

			req.Write(buffer[:n])

			if n < buffer_limit {
				break
			}
		}

		reqS := req.String()

		fmt.Printf("INFO: received \"%s\"\n", reqS)
		res := process(reqS)

		_, err = conn.Write([]byte(res))
		if err != nil {
			fmt.Printf("ERROR: errror during writing data to connection - %s\n", err)
		}

		fmt.Printf("DEBUG: closing connection %s\n", conn.RemoteAddr())
		conn.Close()
	}
}
