package main

import (
	"io"
	"net/http"
	"os"
)

var buff []byte

func main() {
        size := 1024 * 1024 * 1024
        buff = make([]byte, size)
	for i := 0; i < size; i += 1 {
            buff[i] = byte(i)
        }
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	buff[0] = 0
	name, _ := os.Hostname()
	io.WriteString(w, "Hello world from host \"" + name + "\"\n")
}
