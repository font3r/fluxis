package main

import (
	"strings"
)

type CommandName string

const (
	Set    CommandName = "SET"
	Get    CommandName = "GET"
	Delete CommandName = "DELETE"
	Debug  CommandName = "DEBUG"
	Time   CommandName = "TIME"

	key   string = "key"
	value string = "value"
)

var (
	ErrInvalidRequest FluxisError = FluxisError{Code: "INVALID_REQUEST", Message: "Invalid request"}
	ErrInvalidCommand FluxisError = FluxisError{Code: "INVALID_COMMAND", Message: "Invalid command"}
)

var validCommands = []CommandName{Set, Get, Delete, Time, Debug}

type CommandWithArgs struct {
	Command CommandName
	Args    map[string]string
}

func Parse(raw string) (CommandWithArgs, error) {
	parts := strings.Split(strings.Trim(raw, " "), " ")
	if len(parts) == 1 && parts[0] == "" {
		return CommandWithArgs{}, ErrInvalidRequest
	}

	var cmd CommandName
	for _, v := range validCommands {
		if parts[0] == string(v) {
			cmd = v
			break
		}
	}

	switch cmd {
	default:
		fallthrough
	case "":
		return CommandWithArgs{}, ErrInvalidCommand
	case Get:
		return validateGet(parts)
	case Set:
		return validateSet(parts)
	case Delete:
		return validateDelete(parts)
	case Time:
		return validateTime(parts)
	case Debug:
		return validateDebug(parts)
	}
}

func validateGet(parts []string) (CommandWithArgs, error) {
	if len(parts) != 2 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	return CommandWithArgs{
		Command: Get,
		Args: map[string]string{
			key: parts[1],
		},
	}, nil
}

func validateSet(parts []string) (CommandWithArgs, error) {
	if len(parts) != 2 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	kv := strings.Split(parts[1], "=")
	if len(kv) != 2 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	return CommandWithArgs{
		Command: Set,
		Args: map[string]string{
			key:   kv[0],
			value: kv[1],
		},
	}, nil
}

func validateDelete(parts []string) (CommandWithArgs, error) {
	if len(parts) != 2 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	return CommandWithArgs{
		Command: Delete,
		Args: map[string]string{
			key: parts[1],
		},
	}, nil
}

func validateTime(parts []string) (CommandWithArgs, error) {
	if len(parts) != 1 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	return CommandWithArgs{
		Command: Time,
	}, nil
}

func validateDebug(parts []string) (CommandWithArgs, error) {
	if len(parts) != 1 {
		return CommandWithArgs{}, ErrInvalidCommand
	}

	return CommandWithArgs{
		Command: Debug,
	}, nil
}
