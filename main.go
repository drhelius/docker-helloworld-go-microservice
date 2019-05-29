package main

import (
	"io"
	"net/http"
	"os"
        "fmt"
)

var memory []byte

func main() {
        memory = make([]byte, 1024*1024*1024)
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
        memory[0] = 0
	name, _ := os.Hostname()
	io.WriteString(w, "Hello world from host \"" + name + "\"\n")
}
