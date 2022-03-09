package post

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func ParsPayload(args map[string]interface{}) (string, string, error) {
	if _, ok := args["plugin"]; !ok {
		return "", "", fmt.Errorf("missing plugin information")
	}
	plugin := map[string]map[string]string{}
	bytes, ok := args["plugin"].([]byte)
	if !ok {
		return "", "", fmt.Errorf("cast to bytes")
	}
	if err := json.Unmarshal(bytes, &plugin); err != nil {
		return "", "", fmt.Errorf("unmarshal plugin")
	}

	// No we should be able to access the stuff
	inputs, ok := plugin["hello"]
	if !ok {
		return "", "", fmt.Errorf("missing inputs")
	}
	return inputs["channel"], inputs["text"], nil
}

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
