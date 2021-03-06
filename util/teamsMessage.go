package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type TeamsMessage struct {
	Title      string `json:"title"`
	ThemeColor string `json:"themeColor"`
	Text       string `json:"text"`
}

func SendTeamsMessage(title, msg, color, url string) error {
	tBody, _ := json.Marshal(TeamsMessage{
		Title:      title,
		ThemeColor: color,
		Text:       msg,
	})

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(tBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Teams")
	}

	return nil
}
