package ch07

import (
	"io/ioutil"
	"log"
	"net/http"
)

func BuildWebServer() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

func home(response http.ResponseWriter, request *http.Request) {
	content, _ := ioutil.ReadFile("./data/index.html")
	response.Write(content)
}
