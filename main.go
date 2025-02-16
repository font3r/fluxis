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

	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
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

		fmt.Printf("INFO: received %s\n", req.String())

		res := strings.Builder{}
		res.WriteString(req.String())

		_, err = conn.Write([]byte("asd"))
		if err != nil {
			fmt.Printf("ERROR: errror during writing data to connection - %s\n", err)
		}

		fmt.Printf("DEBUG: closing connection %s\n", conn.RemoteAddr())
		conn.Close()
	}
}
