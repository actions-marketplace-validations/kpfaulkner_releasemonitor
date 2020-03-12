package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kpfaulkner/releasemonitor/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func contains(titleList []string, title string) bool {
	for _, t := range titleList {
		if t == title {
			return true
		}
	}
	return false
}

func generateSlackMessage( title string, author string, message string) (models.Webhook, error) {

	fields := []models.Field{
		{
			Title: title,
			Value: message,
			Short: false,
		},
	}

	msg := models.Webhook{
		UserName: "releasemonitor",
		Attachments: []models.Attachment{
			{
				AuthorName: author,
				Fields:     fields,
			},
		},
	}

	return msg, nil
}

func sendToSlack(title string, author string, message string) error {

	msg, err := generateSlackMessage(title, author, message)
	if err != nil {
		fmt.Printf("error generating slack message %s\n", err.Error())
		return err
	}

	endpoint := os.Getenv("SLACK_WEBHOOK")
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "URL is required")
		os.Exit(1)
	}

	enc, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(enc)
	_, err = http.Post(endpoint, "application/json", b)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	dat, err := ioutil.ReadFile(eventPath)
	if err != nil {
    fmt.Printf("unable to read event")
    return
	}

	// have the data, deserialise
  var ev models.ReleaseEventModel
	err = json.Unmarshal(dat, &ev)
	if err != nil {
		fmt.Printf("cannot unmarshal event data")
		return
	}

	checkReleaseName(ev)
}


func checkReleaseName(  ev models.ReleaseEventModel) error {
	if strings.Contains(strings.ToLower(ev.Release.Name), "in test") {
		sendToSlack(ev.Release.Name, ev.Release.Author.Login, "Test release updated")
	}

	return nil
}
