package article

import (
	"encoding/json"
	"os"
)

type Article struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ReadArticles(fName string) (map[string]Article, error) {
	bs, err := os.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	var articles map[string]Article
	err = json.Unmarshal(bs, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
