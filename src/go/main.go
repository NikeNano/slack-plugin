package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	resp := make(map[string]map[string]string)
	resp["node"] = map[string]string{}
	resp["node"]["phase"] = "Succeeded"
	resp["node"]["message"] = "Hello template!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/api/v1/template.execute", hello)
	http.ListenAndServe(":4355", nil)
}
