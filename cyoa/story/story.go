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

func GetStory(fName string) (Story, error) {
	bs, err := os.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	var story Story
	err = json.Unmarshal(bs, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
