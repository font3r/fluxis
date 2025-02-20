package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		rawCmd    string
		expectErr error
		expectRes CommandWithArgs
	}{
		"valid_get": {
			rawCmd: "GET test_key",
			expectRes: CommandWithArgs{
				Command: Get,
				Args:    map[string]string{key: "test_key"},
			},
		},
		"valid_delete": {
			rawCmd: "DELETE test_key",
			expectRes: CommandWithArgs{
				Command: Delete,
				Args:    map[string]string{key: "test_key"},
			},
		},
		"valid_set": {
			rawCmd: "SET key=value",
			expectRes: CommandWithArgs{
				Command: Set,
				Args: map[string]string{
					key:   "key",
					value: "value",
				},
			},
		},
		"valid_debug": {
			rawCmd: "DEBUG",
			expectRes: CommandWithArgs{
				Command: Debug,
			},
		},
		"empty_command": {
			rawCmd:    "",
			expectErr: ErrInvalidRequest,
			expectRes: CommandWithArgs{},
		},
		"unknown_command": {
			rawCmd:    "ASDASD",
			expectErr: ErrInvalidCommand,
			expectRes: CommandWithArgs{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := Parse(test.rawCmd)

			if test.expectErr != err {
				t.Errorf("expect error %s, got %s", test.expectErr, err.Error())
			} else {
				if !reflect.DeepEqual(test.expectRes, res) {
					t.Errorf("expect result %+v, got %+v", test.expectRes, res)
				}
			}
		})
	}

}
