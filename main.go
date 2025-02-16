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

	start(func(s string) string {
		return handleCommand(&st, s)
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
		go processRequest(conn, process)
	}
}

func processRequest(conn net.Conn, cmdHandler func(s string) string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("WARNING: recovered in processRequest() because of:", r)

			internalErr := FluxisError{Code: "INTERNAL_ERROR", Message: "Internal error"}
			_, err := conn.Write([]byte(internalErr.Error()))
			if err != nil {
				fmt.Printf("ERROR: errror during writing data to connection - %s\n", err)
			}

			conn.Close()
		}
	}()

	reqBuilder := strings.Builder{}
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

		reqBuilder.Write(buffer[:n])

		if n < buffer_limit {
			break
		}
	}

	res := cmdHandler(reqBuilder.String())

	_, err := conn.Write([]byte(res))
	if err != nil {
		fmt.Printf("ERROR: errror during writing data to connection - %s\n", err)
	}

	conn.Close()
}

func handleCommand(st *Storage, message string) string {
	fmt.Printf("INFO: raw message %s\n", message)
	cmd, err := Parse(message)

	if err != nil {
		fmt.Printf("ERROR: error during parsing command %s", err.Error())
		return err.Error()
	}

	fmt.Printf("INFO: parsed command %s\n", cmd)

	switch cmd.Command {
	case Set:
		st.SetKey(cmd.Args["KEY"], cmd.Args["VALUE"])
	case Get:
		return st.GetKey(cmd.Args["KEY"]).Value
	case Delete:
		st.DeleteKey(cmd.Args["KEY"])
	case Debug:
		return st.Debug()
	default:
		return FluxisError{Code: "UNSUPPORTED_COMMAND", Message: "Specified command is not supported"}.Error()
	}

	return "OK"
}
