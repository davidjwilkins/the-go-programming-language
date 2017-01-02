//Server1 is a minimal "Echo" server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/favicon.ico", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Println("increment...")
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//counter  echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Println("here")
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
