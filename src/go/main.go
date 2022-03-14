package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/slack-go/slack"
)

func post(channel, text string) error {
	token, ok := os.LookupEnv("SLACK_BOT_TOKEN")
	if !ok {
		return fmt.Errorf("env SLACK_BOT_TOKEN not set")
	}

	api := slack.New(token)
	_, _, err := api.PostMessage(channel, slack.MsgOptionText(text, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return fmt.Errorf("failed to post to slack")
	}
	return nil
}

func parsPayload(args map[string]interface{}) (string, string, error) {
	if _, ok := args["template"]; !ok {
		return "", "", fmt.Errorf("missing template information")
	}
	template, ok := args["template"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	plugin, ok := template["plugin"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	inputs, ok := plugin["test"]
	if !ok {
		return "", "", fmt.Errorf("missing inputs")
	}
	fmt.Println(inputs)
	fmt.Println(reflect.TypeOf(inputs))
	fmt.Println(inputs.(string))
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(inputs.(string)), &sec); err != nil {
		panic(err)
	}
	fmt.Println(sec)
	return sec["channel"].(string), sec["text"].(string), nil //inputs["channel"], inputs["text"], nil
}

func hello(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	args := map[string]interface{}{}
	fmt.Println("requests")
	if err := json.Unmarshal(body, &args); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("parts payload")
	channel, text, err := parsPayload(args)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("LETS POST")
	if err := post(channel, text); err != nil {
		fmt.Println("%w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	fmt.Println("START")
	http.HandleFunc("/api/v1/template.execute", hello)
	http.ListenAndServe(":4355", nil)
}
