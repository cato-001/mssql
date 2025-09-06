package main

import (
	"encoding/json"
	"fmt"
)

func Vjsonln(value any) {
	type valOut struct {
		Status string `json:"status"`
		Value any `json:"value"`
	}
	Jsonln(valOut {
		Status: "success",
		Value: value,
	})
}

func Ejsonln(err error) {
	type errOut struct {
		Status string `json:"status"`
		Error string `json:"error"`
	}
	Jsonln(errOut {
		Status: "error",
		Error: err.Error(),
	})
}

func Jsonln(out any) {
	content, err := json.Marshal(out)
	if err == nil {
		fmt.Println(string(content))
	}
}
