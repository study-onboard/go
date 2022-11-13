package ch05

import (
	"flag"
	"fmt"
)

func AnonymousFunction() {
	// define to call
	func(message string) {
		fmt.Printf("message: %s\n", message)
	}("Show me the money")

	// use as variable 1
	var handler1 func(string, string) = func(requester string, job string) {
		fmt.Printf("handling for %s: %s\n", requester, job)
	}
	handler1("Kut Zhang", "Deploy to Prod Cluster")

	// use as variable 2
	var handler2 func(string, string) (bool, string) = func(requester string, job string) (success bool, description string) {
		success = true
		description = fmt.Sprintf("%s request job `%s` complete", requester, job)
		return
	}
	success, description := handler2("Lana Chang", "QA")
	if success {
		fmt.Printf("job handled: %s\n", description)
	} else {
		fmt.Printf("job handle failed: %s\n", description)
	}

	// use as parameter
	processJob("BuildProject", func(result string) {
		fmt.Printf("Build project result: %s\n", result)
	})

	// using function as map value
	usingFunctionAsMapValue()
}

func usingFunctionAsMapValue() {
	var ptr *string = flag.String("command", "default", "command to be executed")
	flag.Parse()

	var commands = map[string]func(){
		"build": func() {
			fmt.Println("Building....................................")
		},

		"package": func() {
			fmt.Println("Packing.....................................")
		},

		"default": func() {
			fmt.Println("Nothing to do for default command")
		},
	}
	var command = *ptr
	var handler = commands[command]
	if handler != nil {
		handler()
	} else {
		fmt.Printf("Can't find handler for %s\n", command)
	}
}

func processJob(jobName string, callback func(string)) {
	fmt.Printf("Procing job: %s\n", jobName)
	fmt.Printf("Job `%s` complete.\n", jobName)
	callback("complete")
}
