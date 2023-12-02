package urlshortener

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, mux *http.ServeMux) *http.ServeMux {
	for k, v := range pathsToUrls {
		mux.Handle(k, http.RedirectHandler(v, http.StatusSeeOther))
	}
	return mux
}

type UrlShorten struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, mux *http.ServeMux) (*http.ServeMux, error) {
	var urlshortens []UrlShorten
	err := yaml.Unmarshal(yml, &urlshortens)
	if err != nil {
		return mux, err
	}
	for _, v := range urlshortens {
		mux.Handle(v.Path, http.RedirectHandler(v.Url, http.StatusSeeOther))
	}
	return mux, nil
}
