package ch04

import "fmt"

func UsingCondition() {
	if content, error := requestSapi(true); error != nil {
		fmt.Printf("request sapi error: %s\n", *error)
	} else {
		fmt.Printf("request sapi success: %s\n", *content)
	}
	if content, error := requestSapi(false); error != nil {
		fmt.Printf("request sapi error: %s\n", *error)
	} else {
		fmt.Printf("request sapi success: %s\n", *content)
	}

}

func requestSapi(t bool) (*string, *string) {
	if t {
		content := "show me the money"
		return &content, nil
	} else {
		error := "request failed"
		return nil, &error
	}
}
