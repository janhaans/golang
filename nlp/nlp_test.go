package nlp

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

type TestCases struct {
	Cases []TokenCase
}
type TokenCase struct {
	Text   string
	Tokens []string
}

func TestTokenizeToml(t *testing.T) {
	var tcs TestCases
	_, err := toml.DecodeFile("tokenize_cases.toml", &tcs)
	require.NoError(t, err, fmt.Sprintf("%#v", err))
	for _, tc := range tcs.Cases {
		t.Run(tc.Text, func(t *testing.T) {
			require.Equal(t, tc.Tokens, Tokenize(tc.Text))
		})
	}
}

var tokenCases = []struct {
	text   string
	tokens []string
}{
	{"Who's on first", []string{"who", "s", "on", "first"}},
	{"Who's on second", []string{"who", "s", "on", "second"}},
	{"", nil},
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range tokenCases {
		t.Run(tc.text, func(t *testing.T) {
			require.Equal(t, tc.tokens, Tokenize(tc.text))
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "Who's on second?"
	expected := []string{"who", "s", "on", "second"}
	got := Tokenize(text)
	require.Equal(t, expected, got)
	/*
		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("Expected %#v, got %#v\n", expected, got)
		}
	*/
}
