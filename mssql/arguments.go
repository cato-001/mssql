package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func Arguments(args string) ([]any, error) {
	var jsonArgs any
	err := json.Unmarshal([]byte(args), &jsonArgs)
	if err != nil {
		return nil, fmt.Errorf("while parsing json arguments: %e", err)
	}

	arguments := make([]any, 0)
	switch argsValue := jsonArgs.(type) {
	case []any:
		arguments = argsValue
	case map[string]any:
		for key, value := range argsValue {
			arguments = append(arguments, sql.Named(key, value))
		}
	}
	return arguments, nil
}

