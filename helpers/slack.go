package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/pigeon/settings"
)

func SendSlackMessage(msg string) {

	type TextBlock struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
	type MainBlock struct {
		Type string    `json:"type"`
		Text TextBlock `json:"text"`
	}
	type ResponseBlock struct {
		Blocks []MainBlock `json:"blocks"`
	}

	//downMsg := fmt.Sprintf("*%s Is Down* :crying_cat_face:\n Blind Cat cannot reach the service and marked it as down.\n*Total missing cats: %d*", serviceName, errorCount)
	//upMsg := fmt.Sprintf("*%s Is Up!* :smile_cat:\n Blind Cat marked the service up after a hectic downtime.", serviceName)
	//sendMsg := upMsg
	//
	textBlock := TextBlock{
		Type: "mrkdwn",
		Text: msg,
	}

	r := MainBlock{Type: "section", Text: textBlock}

	slackMsg := ResponseBlock{Blocks: []MainBlock{r}}

	jsonStr, err := json.Marshal(slackMsg)

	hookUrl := settings.SystemAppConfig.SlackUrl
	resp, err := requests.PostJson(hookUrl, string(jsonStr))
	fmt.Println(string(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
