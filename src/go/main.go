package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("body: ", body)
	resp := make(map[string]map[string]string)
	resp["node"] = map[string]string{}
	resp["node"]["phase"] = "Succeeded"
	resp["node"]["message"] = "Hello template!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	time.Sleep(time.Minute * 10)
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/api/v1/template.execute", hello)
	http.ListenAndServe(":4355", nil)
}
