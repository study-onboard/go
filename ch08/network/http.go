package network

import "fmt"

func Request(url string, method string, parameters *map[string]any) (string, error) {
	fmt.Printf("network request - %s(%s) - %s\n", url, method, parameters)
	return "success", nil
}

func CheckHealth(serverAddress string, serverPort int) bool {
	fmt.Printf("network check health - %s:%d\n", serverAddress, serverPort)
	return true
}
