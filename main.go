package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	name, _ := os.Hostname()
	io.WriteString(w, "Hello world from host \"" + name + "\"\n")
}
