package ch07

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func BuildWebServer() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func home(response http.ResponseWriter, request *http.Request) {
	content, _ := os.ReadFile("./data/index.html")
	_, err := response.Write(content)
	if err != nil {
		fmt.Printf("Can't write any content to the response: %s", err.Error())
	}
}
