package ch04

import (
	"flag"
	"fmt"
	"strconv"
	"unsafe"
)

func SwitchStatement() {
	context := prepare()
	simpleUse(context)
	conditionCaseUse(context)
}

func prepare() map[string]*string {
	var context = make(map[string]*string, 0)

	var ptr *string = flag.String("env", "default", "application *environment, such as dev / sit / uat / prod")
	context["env"] = ptr

	ptr = flag.String("age", "6", "human age, such as 6, 12, 24, 36, 54....etc.")
	context["age"] = ptr

	flag.Parse()
	return context
}

func simpleUse(context map[string]*string) {
	var envPtr *string = (*string)(unsafe.Pointer(context["env"]))

	switch *envPtr {
	case "dev", "sit":
		fmt.Printf("Deploying to local cluster for: %s\n", *envPtr)

	case "uat":
		fmt.Printf("Deploying to QA cluster for `%s`, and notification had been sent to QA staff\n", *envPtr)

	case "prod":
		fmt.Printf("Deploying to PROD cluster for `%s`, and notificaton has been sent to product commander\n", *envPtr)

	default:
		fmt.Printf("Nothing to for `%s`\n", *envPtr)
	}
}

func conditionCaseUse(context map[string]*string) {
	agePtr := context["age"]

	age, error := strconv.Atoi(*agePtr)
	if error == nil {
		switch {
		case age < 10:
			fmt.Println("Is a child")

		case age < 20 && age > 10:
			fmt.Println("Is a juvenile")

		case age > 20 && age < 45:
			fmt.Println("May be a good age")
			fallthrough

		default:
			fmt.Println("Is a strong man")
		}
	} else {
		fmt.Printf("Invalid age format: %s\n", *agePtr)
	}
}
