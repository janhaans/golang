package story

import (
	"encoding/json"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Path string `json:"arc"`
}

var PublishedStory Story

func PublishStory(filename string) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, &PublishedStory)
	if err != nil {
		return err
	}

	return nil
}
