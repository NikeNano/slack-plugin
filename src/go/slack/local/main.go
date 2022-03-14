package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	fmt.Println(1)
	if _, ok := args["template"]; !ok {
		return "", "", fmt.Errorf("missing template information")
	}
	template, ok := args["template"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	fmt.Println(2)
	plugin, ok := template["plugin"].(map[string]string)
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	fmt.Println(3)
	inputs, ok := plugin["test"]
	if !ok {
		return "", "", fmt.Errorf("missing inputs")
	}
	fmt.Println(4)
	var info map[string]interface{}
	json.Unmarshal([]byte(inputs), &info)

	fmt.Println("Success")
	fmt.Println(info)
	return info["channel"].(string), info["text"].(string), nil //inputs["channel"], inputs["text"], nil
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
	// fmt.Println("START")
	// http.HandleFunc("/api/v1/template.execute", hello)
	// http.ListenAndServe(":4355", nil)
	err := post("C035Q8CELGM", "hello world")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
