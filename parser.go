package main

import (
	"fmt"
	"strings"
)

type CommandName string

const (
	Set   CommandName = "SET"
	Get   CommandName = "GET"
	Debug CommandName = "DEBUG"
)

var validCommands = []CommandName{Set, Get, Debug}

type CommandWithArgs struct {
	Command CommandName
	Args    map[string]string
}

/*
	SET key=value
	GET key
	LIST
*/

func Parse(raw string) (CommandWithArgs, error) {
	for _, v := range validCommands {
		if !strings.HasPrefix(raw, string(v)) {
			continue
		}

		args := strings.Split(
			strings.TrimLeft(
				raw,
				fmt.Sprintf("%s ", string(v))),
			"=")

		argsMap := make(map[string]string)
		argsC := len(args)
		if argsC == 0 {
			argsMap["KEY"] = ""
		}

		if argsC >= 1 {
			argsMap["KEY"] = args[0]
		}

		if argsC >= 2 {
			argsMap["VALUE"] = args[1]
		}

		return CommandWithArgs{
			Command: v,
			Args:    argsMap,
		}, nil
	}

	return CommandWithArgs{}, FluxisError{Code: "INVALID_COMMAND", Message: "Invalid command"}
}
