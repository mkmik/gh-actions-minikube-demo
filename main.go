package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "listen address")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
