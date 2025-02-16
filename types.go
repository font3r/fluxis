package main

import "fmt"

type FluxisError struct {
	Code    string
	Message string
}

func (fe FluxisError) Error() string {
	return fmt.Sprintf("%s - %s", fe.Code, fe.Message)
}
