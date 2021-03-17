package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(1)
		time.Sleep(time.Hour)
	})
	http.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(2)
		time.Sleep(time.Hour)
	})

	runtime.GOMAXPROCS(1)
	http.ListenAndServe(":8080", nil)
}
