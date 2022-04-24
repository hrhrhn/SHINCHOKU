package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func index_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<DEFAULT> Hello World, %s!", request.URL.Path[1:])
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world/", log(world))

	server.ListenAndServe()
}
