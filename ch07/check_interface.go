package ch07

import (
	"errors"
	"flag"
	"fmt"
)

func CheckInterface() {
	context := prepare()

	case1(context)
	fmt.Println("-------------------------------------------------------------------------")

	case2()
	fmt.Println("-------------------------------------------------------------------------")

	case3(context)
}

// ---------------------------------------------------------------------------------------------

func case1(context map[string]any) {
	mode, ok := context["mode"].(*string)
	if ok {
		fmt.Printf("Application runtime mode: %s\n", *mode)
	} else {
		fmt.Println("Invalid mode!")
	}

	times, ok := context["times"].(*int)
	if ok {
		fmt.Printf("Application run times: %d\n", *times)
	} else {
		fmt.Println("Invalid times!")
	}
}

func prepare() map[string]any {
	var context = make(map[string]any, 0)

	var ptr any = flag.String("mode", "test", "application mode, such as test / qa / prod")
	context["mode"] = ptr

	ptr = flag.Int("times", 6, "application run times, such as 6, 12, 24, 36, 54....etc.")
	context["times"] = ptr

	flag.Parse()
	return context
}

// ---------------------------------------------------------------------------------------------

func case2() {
	var errs []error
	errs = append(errs, &NetworkError{})
	errs = append(errs, &ConfigError{})
	errs = append(errs, errors.New("Common error"))

	for _, err := range errs {
		switch err.(type) {
		case *NetworkError:
			fmt.Printf("Network error: %s\n", err.Error())

		case *ConfigError:
			fmt.Printf("Config error: %s\n", err.Error())

		default:
			fmt.Printf("Common error: %s\n", err.Error())
		}
	}
}

type NetworkError struct {
}

func (err *NetworkError) Error() string {
	return "Could not connect to server."
}

type ConfigError struct {
}

func (err *ConfigError) Error() string {
	return "Could not read configs from server."
}

// ---------------------------------------------------------------------------------------------

func case3(context map[string]any) {
	modePtr := context["mode"].(*string)
	mode := *modePtr
	fmt.Printf("Mode: %s\n", mode)

	timesPtr := context["times"].(*int)
	times := *timesPtr
	fmt.Printf("Times: %d\n", times)
}
