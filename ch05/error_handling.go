package ch05

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// references
// --------------------------------------------------------------
// https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go

func ErrorHandling() {
	turnOnNetwork()
	result, error := requestService("platform-sapi", "Kut Zhang")
	if error != nil {
		fmt.Printf("request service fail: %s\n", error.Error())
	} else {
		fmt.Printf("request service result: %s\n", result)
	}

	turnOffNetwork()
	result, error = requestService("platform-sapi", "Kut Zhang")
	if error != nil {
		fmt.Printf("request service fail: %s\n", error.Error())
	} else {
		fmt.Printf("request service result: %s\n", result)
	}

	testCustomError()
}

var networkHealth bool = true

func turnOnNetwork() {
	networkHealth = true
}

func turnOffNetwork() {
	networkHealth = false
}

func requestService(service string, requester string) (string, error) {
	if networkHealth {
		return "success", nil
	} else {
		return "", errors.New("Network unreachable")
	}
}

type serviceRequestError struct {
	message string
	traceId string
}

func (error *serviceRequestError) Error() string {
	return error.message
}

func makeError() (string, error) {
	return "", &serviceRequestError{
		message: "Make war, no love",
		traceId: uuid.NewString(),
	}
}

func testCustomError() {
	result, error := makeError()
	if error != nil {
		err, _ := error.(*serviceRequestError)
		fmt.Printf("catch custom error: %s, %s\n", err.traceId, err.message)
	} else {
		fmt.Printf("result: %s\n", result)
	}

	switch xx := error.(type) {
	case *serviceRequestError:
		fmt.Printf("catch custom error: %s, %s\n", xx.traceId, xx.message)
	default:
		fmt.Printf("catch default error: %s\n", xx.Error())
	}
}
