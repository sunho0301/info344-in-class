package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	//w.Write([]byte("Hello, World!"))
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)

}

func memoryHandler(w http.ResponseWriter, r *http.Request) {

	runtime.GC()

	// & means we want pointer not the instance
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)

}

func main() {
	//fmt.Println("Hello, World!")
	mux := http.NewServeMux()
	// if "/hello" <- must be "hello". "/hello/" includes "/hello/foo"
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)
	fmt.Printf("server is listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))

}
