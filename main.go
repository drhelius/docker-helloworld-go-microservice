package main

import (
	"io"
	"net/http"
	"os"
)

var buff []byte

func main() {
        buff = make([]byte, 1024*1024*1024, 1024*1024*1024)
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
        size := 1024 * 1024 * 1024
	for i := 0; i < size; i += 1 {
            buff[i] = byte(i)
        }
	name, _ := os.Hostname()
	io.WriteString(w, "Hello world from host \"" + name + "\"\n")
}
