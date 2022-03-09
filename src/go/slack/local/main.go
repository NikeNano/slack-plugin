package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func Post(channel, text string) error {
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

func ParsPayload(args map[string]interface{}) (string, string, error) {
	if _, ok := args["plugin"]; !ok {
		return "", "", fmt.Errorf("missing plugin information")
	}
	plugin, ok := args["plugin"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	inputs, ok := plugin["hello"]
	if !ok {
		return "", "", fmt.Errorf("missing inputs")
	}
	info, ok := inputs.(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("failed to parse plugin")
	}

	return info["channel"].(string), info["text"].(string), nil //inputs["channel"], inputs["text"], nil
}

func main() {
	out := map[string]interface{}{
		"workflow": map[string]interface{}{"metadata": map[string]interface{}{"name": "output-parameter-8pmnt"}},
		"template": map[string]interface{}{"name": "slack-integration", "inputs": map[string]interface{}{"parameters": []map[string]interface{}{{"name": "message", "value": map[string]interface{}{"channel": "test", "text": "Hello Niklas"}}}}},
		"outputs":  "",
		"metadata": "",
		"plugin":   map[string]interface{}{"hello": map[string]interface{}{"channel": "test", "text": "Hello Niklas"}},
	}
	channel, text, err := ParsPayload(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(channel, text)
	if err := Post("C035Q8CELGM", "hello"); err != nil {
		panic(err)
	}

}
