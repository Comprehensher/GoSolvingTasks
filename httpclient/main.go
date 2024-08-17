package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	response, err := http.Get("http://localhost:5000/html")
	if err == nil {
		response.Write(os.Stdout)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
