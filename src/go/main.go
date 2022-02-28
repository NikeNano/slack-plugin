package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Success")
}

func main() {
	http.HandleFunc("/api/v1/template.execute", hello)
	http.ListenAndServe(":4355", nil)
}
